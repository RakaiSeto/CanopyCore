package main

import (
	"canopyCore/modules"
	"context"
	"database/sql"
	"fmt"
	"runtime"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

var db *sql.DB
var rc *redis.Client
var cx context.Context
var connRabbit *amqp.Connection
var chIncoming *amqp.Channel

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
}