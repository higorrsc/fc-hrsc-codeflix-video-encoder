package main

import (
	"log"
	"os"
	"strconv"

	"github.com/higorrsc/fc-hrsc-codeflix-video-encoder/application/services"
	"github.com/higorrsc/fc-hrsc-codeflix-video-encoder/framework/database"
	"github.com/higorrsc/fc-hrsc-codeflix-video-encoder/framework/queue"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

var db database.Database

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	autoMigrateDb, err := strconv.ParseBool(os.Getenv("AUTO_MIGRATE_DB"))
	if err != nil {
		log.Fatalf("Error parsing AUTO_MIGRATE_DB: %v", err)
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatalf("Error parsing DEBUG: %v", err)
	}

	db.AutoMigrateDb = autoMigrateDb
	db.Debug = debug
	db.Dsn = os.Getenv("DSN")
	db.DbType = os.Getenv("DB_TYPE")
	db.DsnTest = os.Getenv("DSN_TEST")
	db.DbTypeTest = os.Getenv("DB_TYPE_TEST")
	db.Env = os.Getenv("ENV")
}

func main() {
	messageChannel := make(chan amqp.Delivery)
	jobReturnChannel := make(chan services.JobWorkerResult)

	dbConnection, err := db.Connect()
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}
	defer dbConnection.Close()

	rabbitMQ := queue.NewRabbitMQ()
	ch := rabbitMQ.Connect()
	defer ch.Close()

	rabbitMQ.Consume(messageChannel)

	jobManager := services.NewJobManager(dbConnection, rabbitMQ, jobReturnChannel, messageChannel)
	jobManager.Start(ch)
}
