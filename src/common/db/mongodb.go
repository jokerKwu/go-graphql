package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	UserCollection *mongo.Collection
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

func InitMongo() error {
	clientOptions := options.Client()
	clientOptions = clientOptions.ApplyURI("mongodb+srv://ryan:tngkr485@cluster0.ke0pv.mongodb.net/?retryWrites=true&w=majority")
	clientOptions.SetMaxPoolSize(1)
	clientOptions.SetMinPoolSize(1)
	clientOptions.SetMaxConnIdleTime(20 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	MongoClient, _ = mongo.Connect(ctx, clientOptions)
	if err := PingMongo(); err != nil {
		return err
	}
	if err := InitCollection(); err != nil {
		return err
	}
	return nil
}

func InitCollection() error {
	mongo := MongoClient.Database("medical_dev")
	UserCollection = mongo.Collection("user")
	return nil
}

func PingMongo() error {
	err := MongoClient.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}
	fmt.Println("db 핑 통과")
	return nil
}
