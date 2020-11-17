package mongodb

import (
	"autoposter/config"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
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
func (mongo *MongoDataStore) CreateDataStore(dbName string, connectionStr string) {
	db, client := connect(dbName, connectionStr)

	if db == nil && client == nil {
		logrus.Fatal("Failed to connect to database")
	}

	mongo.db = db
	mongo.Client = client
}

func connect(dbName string, connectionStr string) (*mongo.Database, *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var client *mongo.Client

	connectOnce.Do(func() {
		db, client = connectToMongo(dbName, connectionStr)
	})

	return db, client
}

func connectToMongo(dbName string, connectionStr string) (*mongo.Database, *mongo.Client) {
	cfg := config.GetConfig()
	var err error

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionStr))
	if err != nil {
		logrus.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logrus.Fatal(err)
	}

	var db = client.Database(dbName)
	fmt.Printf("Connected to DbName: %s\n", cfg.DbName)

	return db, client
}
