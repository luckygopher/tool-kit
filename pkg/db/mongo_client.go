package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

func InitMongoClient(cfg MongoConfig) error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), cfg.GetDefaultTimeout())
	defer cancel()
	ops := options.Client()
	ops.ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s/?authSource=%s", cfg.UserName, cfg.PassWord, cfg.Host, cfg.AuthSource))
	ops.SetMaxPoolSize(cfg.MaxPoolSize)
	MongoClient, err = mongo.Connect(ctx, ops)
	if err != nil {
		return err
	}
	if err := MongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	return nil
}
