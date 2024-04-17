package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCollection interface {
	// QueryList 查询列表
	QueryList(ctx context.Context, filter interface{}, sort interface{}) ([]interface{}, error)
	// QueryOne 查询单条
	QueryOne(ctx context.Context, filter interface{}) (interface{}, error)
	// InsertOne 插入单条
	InsertOne(ctx context.Context, model interface{}) (interface{}, error)
	// InsertMany 插入多条
	InsertMany(ctx context.Context, models []interface{}) ([]interface{}, error)
}

func NewMongoCollection(database, collection string) MongoCollection {
	return MongoCollectionImp{Collection: MongoClient.Database(database).Collection(collection)}
}

type MongoCollectionImp struct {
	Collection *mongo.Collection
}

func (m MongoCollectionImp) QueryList(ctx context.Context, filter interface{}, sort interface{}) ([]interface{}, error) {
	var err error

	opts := options.Find().SetSort(sort)
	finder, err := m.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, 0)
	if err := finder.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, err
}

func (m MongoCollectionImp) QueryOne(ctx context.Context, filter interface{}) (interface{}, error) {
	result := new(interface{})
	err := m.Collection.FindOne(ctx, filter, options.FindOne()).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m MongoCollectionImp) InsertOne(ctx context.Context, model interface{}) (interface{}, error) {
	result, err := m.Collection.InsertOne(ctx, model, options.InsertOne())
	if err != nil {
		return nil, err
	}
	return result.InsertedID, err
}

func (m MongoCollectionImp) InsertMany(ctx context.Context, models []interface{}) ([]interface{}, error) {
	result, err := m.Collection.InsertMany(ctx, models, options.InsertMany())
	if err != nil {
		return nil, err
	}
	return result.InsertedIDs, err
}
