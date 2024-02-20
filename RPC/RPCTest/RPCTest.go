package main

import (
	"canopyCore/modules"
	"context"
	"database/sql"
	"fmt"
	"net"
	"runtime"
	"time"

	test "canopyCore/grpc/test"

	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

var db *sql.DB
var rc *redis.Client
var cx context.Context
var connRabbit *amqp.Connection
var chIncoming *amqp.Channel
var srv *grpc.Server

const THEPORT = "20000"

type TestServer struct{
	test.HelloWorldServiceServer
}

var testServer TestServer

func (testServer TestServer) DoHelloWorld(ctx context.Context, request *test.EmptyRequest) (*test.HelloWorldResponse, error) {
	localResponse := new(test.HelloWorldResponse)
	localResponse.Hello = "This Message is from GRPC"

	return localResponse, nil
}

func main() {
	// Load configuration file
	modules.InitiateGlobalVariables()
	runtime.GOMAXPROCS(4)

	// Initiate Database
	var errDB error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		modules.MapConfig["pgsqlHost"], modules.MapConfig["pgsqlPort"], modules.MapConfig["pgsqlUser"],
		modules.MapConfig["pgsqlPassword"], modules.MapConfig["pgsqlName"])
	db, errDB = sql.Open("postgres", psqlInfo) // db udah di defined diatas, jadi harus pake = bukan :=

	if errDB != nil {
		modules.Logging(modules.Resource(), "STARTING UP", "START SERVICE", "SERVER IP", "Postgres unconnected", errDB)
		panic(errDB)
	} else {
		modules.Logging(modules.Resource(), "STARTING UP", "START SERVICE", "SERVER IP", "Postgres connected", nil)
	}

	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(50)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			modules.Logging(modules.Resource(), "STARTING UP", "START SERVICE", "SERVER IP", "Failed to close Postgres", err)
		} else {
			modules.Logging(modules.Resource(), "STARTING UP", "START SERVICE", "SERVER IP", "Success to close Postgres", nil)
		}
	}(db)

	errDB = db.Ping()
	if errDB != nil {
		panic(errDB)
	}

	// Initiate Rabbit
	var errRabbit error
	connRabbit, errRabbit = amqp.Dial("amqp://" +
		modules.MapConfig["rabbitMqUser"] + ":" +
		modules.MapConfig["rabbitMqPassword"] + "@" +
		modules.MapConfig["rabbitMqHost"] + ":" +
		modules.MapConfig["rabbitMqPort"] + "/" +
		modules.MapConfig["rabbitMqVHost"])
	if errRabbit != nil {
		modules.Logging(modules.Resource(), "STARTING UP", "START SERVICE", "SERVER IP", "RabbitMQ unconnected", errRabbit)
		panic(errRabbit)
	} else {
		modules.Logging(modules.Resource(), "STARTING UP", "START SERVICE", "SERVER IP", "RabbitMQ connected", nil)
	}
	//defer connRabbit.Close()

	var errT error
	chIncoming, errT = connRabbit.Channel()
	if errT != nil {
		panic(errT)
	}

	// Initiate Redis
	rc = modules.InitiateRedisClient()
	cx = context.Background()
	errRedis := rc.Ping(cx).Err()
	if errRedis != nil {
		modules.Logging(modules.Resource(), "STARTING UP", "START SERVICE", "SERVER IP", "Redis unconnected", errRedis)
		panic(errRedis)
	} else {
		modules.Logging(modules.Resource(), "STARTING UP", "START SERVICE", "SERVER IP", "Redis connected", nil)
	}

	customFunc := func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(customFunc),
	}

	opt1 := grpc.MaxConcurrentStreams(250)
	srv = grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_recovery.UnaryServerInterceptor(opts...),
	),
	grpc_middleware.WithStreamServerChain(
		grpc_recovery.StreamServerInterceptor(opts...),
	),opt1)

	test.RegisterHelloWorldServiceServer(srv, testServer)
	modules.Logging(modules.Resource(), "STARTING RPC TESTING", "SERVER", "SERVER IP", "RPC TESTING loaded", nil)

	l, err := net.Listen("tcp", ":"+THEPORT)

	if err != nil {
		modules.Logging(modules.Resource(), "STARTING RPC TESTING", "START RPC", "SERVER IP", "RPC Failed to start and listen port : "+THEPORT, err)
	} else {
		modules.Logging(modules.Resource(), "STARTING RPC TESTING", "START RPC", "SERVER IP", "RPC Start and Listen on port : "+THEPORT, nil)
		fmt.Println("SERVICE RPC TESTING IS ACTIVE ON PORT " + THEPORT)
	}

	log.Fatal(srv.Serve(l))
}