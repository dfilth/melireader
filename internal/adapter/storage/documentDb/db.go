package documentDb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"melireader/internal/adapter/config"
	_ "os"
	"time"
)

var meliDb *mongo.Database

func DocumentDB(connStr, database string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(connStr)
	print(connStr)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client.Database(database), nil
}

func MeliDB(dbConfig *config.DB) *mongo.Database {
	if meliDb != nil {
		return meliDb
	}

	bd, err := DocumentDB(dbConfig.Connection, dbConfig.Name)
	if err != nil {
		fmt.Printf("Error getting connection with mongo: (%v)", dbConfig.Connection)
		panic(err)
	}

	meliDb = bd

	return meliDb
}
