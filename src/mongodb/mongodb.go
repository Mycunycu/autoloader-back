package mongodb

import (
	"autoposter/config"
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDataStore - ...
type MongoDataStore struct {
	db     *mongo.Database
	Client *mongo.Client
}

// CreateDataStore - ...
func CreateDataStore() *MongoDataStore {
	var mongoDataStore *MongoDataStore
	db, client := connect()

	if db == nil && client == nil {
		log.Fatal("Failed to connect to database")
	}

	mongoDataStore = new(MongoDataStore)
	mongoDataStore.db = db
	mongoDataStore.Client = client

	return mongoDataStore
}

func connect() (*mongo.Database, *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var client *mongo.Client

	connectOnce.Do(func() {
		db, client = connectToMongo()
	})

	return db, client
}

func connectToMongo() (*mongo.Database, *mongo.Client) {
	cfg := config.GetConfig()
	var err error

	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.MongoDbURL))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var db = client.Database(cfg.DbName)
	fmt.Printf("Connected to DbName: %s\n", cfg.DbName)

	return db, client
}
