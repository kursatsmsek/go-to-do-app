package configuration

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfiguration struct {
	client   *mongo.Client
	database *mongo.Database
}

func (mc MongoConfiguration) Init(uri string, dbName string) MongoConfiguration {
	client := createMongoClient(uri)
	database := client.Database(dbName)
	return MongoConfiguration{
		client:   client,
		database: database,
	}
}

func (mc MongoConfiguration) Database() *mongo.Database {
	return mc.database
}

func createMongoClient(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}
	return client
}
