package main

import (
	"CanopyCore/RPC/RPCCall"
	rootadminweb "CanopyCore/grpc/rootadminweb"
	"CanopyCore/modules"
	"context"
	"database/sql"
	"fmt"
	"net"
	"runtime"
	"time"

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

const THEPORT = "20020"

type RootAdminWebServer struct {
	rootadminweb.RootAdminWebServiceServer
}

var rootAdminWebServer RootAdminWebServer

func (rootAdminWebServer RootAdminWebServer) DoLogin(ctx context.Context, request *rootadminweb.DoLoginRequest) (*rootadminweb.DoLoginResponse, error) {
	localResponse, _ := RPCCall.CallFunctionDoLoginRootAdminWeb(ctx, db, rc, request)
	return localResponse, nil
}

func (rootAdminWebServer RootAdminWebServer) DoLogout(ctx context.Context, request *rootadminweb.DoLogoutRequest) (*rootadminweb.DoLogoutResponse, error) {
	localResponse, _ := RPCCall.CallFunctionDoLogoutRootAdminWeb(ctx, db, rc, request)
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
		modules.Logging(modules.Resource(), "STARTING RPC ROOT ADMIN WEB", "SERVER", "SERVER IP", "Postgres unconnected", errDB)
		panic(errDB)
	} else {
		modules.Logging(modules.Resource(), "STARTING RPC ROOT ADMIN WEB", "SERVER", "SERVER IP", "Postgres connected", nil)
	}

	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(50)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			modules.Logging(modules.Resource(), "STARTING RPC ROOT ADMIN WEB", "SERVER", "SERVER IP", "Failed to close Postgres", err)
		} else {
			modules.Logging(modules.Resource(), "STARTING RPC ROOT ADMIN WEB", "SERVER", "SERVER IP", "Success to close Postgres", nil)
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
		modules.Logging(modules.Resource(), "STARTING RPC ROOT ADMIN WEB", "SERVER", "SERVER IP", "RabbitMQ unconnected", errRabbit)
		panic(errRabbit)
	} else {
		modules.Logging(modules.Resource(), "STARTING RPC ROOT ADMIN WEB", "SERVER", "SERVER IP", "RabbitMQ connected", nil)
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
		modules.Logging(modules.Resource(), "STARTING RPC ROOT ADMIN WEB", "SERVER", "SERVER IP", "Redis unconnected", errRedis)
		panic(errRedis)
	} else {
		modules.Logging(modules.Resource(), "STARTING RPC ROOT ADMIN WEB", "SERVER", "SERVER IP", "Redis connected", nil)
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
		), opt1)

	rootadminweb.RegisterRootAdminWebServiceServer(srv, rootAdminWebServer)
	modules.Logging(modules.Resource(), "STARTING RPC ROOT ADMIN WEB", "SERVER", "SERVER IP", "RPC ROOT ADMIN WEB loaded", nil)

	l, err := net.Listen("tcp", ":"+THEPORT)

	status := ""
	if err != nil {
		status = "FAILED"
		modules.Logging(modules.Resource(), "STARTING RPC ROOT ADMIN WEB", "SERVER", "SERVER IP", "RPC Failed to start and listen port : "+THEPORT, err)
	} else {
		status = "SUCCESS"
		modules.Logging(modules.Resource(), "STARTING RPC ROOT ADMIN WEB", "SERVER", "SERVER IP", "RPC Start and Listen on port : "+THEPORT, nil)
		fmt.Println("SERVICE RPC ROOT ADMIN WEB IS ACTIVE ON PORT " + THEPORT)
	}

	modules.SaveWebActivity(db, "STARTING RPC ROOT ADMIN WEB", modules.Resource(), "SERVER", "SERVER IP", modules.DoFormatDateTime("YYYY-0M-0D HH:mm:ss.S", time.Now()), "STARTING RPC ROOT ADMIN WEB", status)

	log.Fatal(srv.Serve(l))
}
