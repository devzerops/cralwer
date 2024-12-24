
package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func InitMongoDB(uri string) error {
	var err error
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	return mongoClient.Ping(context.TODO(), nil)
}

func GetMongoClient() *mongo.Client {
	return mongoClient
}

func CloseMongoDB() error {
	return mongoClient.Disconnect(context.TODO())
}