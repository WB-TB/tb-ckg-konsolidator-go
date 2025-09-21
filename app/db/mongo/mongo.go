package mongo

import (
	"context"
	"fhir-sirs/app/config"
	"log"
	"sync"
	"time"

	mongoDB "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	dbConn *mongoDB.Client
	lockDB sync.Mutex
)

func GetConnection() *mongoDB.Client {
	if dbConn == nil {
		lockDB.Lock()
		defer lockDB.Unlock()
		dbConn = ConnectDB()
	}
	return dbConn
}

func NewConnectionDB() (*mongoDB.Client, error) {
	dbConn = ConnectDB()
	return dbConn, nil
}

func ConnectDB() *mongoDB.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongoDB.Connect(
		ctx,
		options.Client().SetRetryWrites(true),
		options.Client().SetRetryReads(true),
		options.Client().SetMinPoolSize(5),
		options.Client().SetMaxConnecting(100),
		options.Client().ApplyURI(config.GetConfig().MongoConnectionString),
	)
	if err != nil {
		log.Print(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Panicf("got an error while connecting database server, error: %s", err)
	}

	return client
}
