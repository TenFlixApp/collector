package data

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectToDB() {
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_CONN_STRING"))
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
}

func FindData(collectionName string, filter bson.D) (results []bson.M, err error) {
	cursor, err := client.Database("metrics").Collection(collectionName).Find(context.Background(), filter)
	if err != nil {
		log.Println("Error in Find: ", err)
		return nil, err
	}
	if err = cursor.All(context.Background(), &results); err != nil {
		log.Println("Error in cursor.All: ", err)
		return nil, err
	}
	return results, nil
}

func AggregateData(collectionName string, pipeline bson.A) (results []bson.M, err error) {
	cursor, err := client.Database("metrics").Collection(collectionName).Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Println("Error in Aggregate: ", err)
		return nil, err
	}
	if err = cursor.All(context.Background(), &results); err != nil {
		log.Println("Error in cursor.All: ", err)
		return nil, err
	}
	return results, nil
}

func InsertData(collectionName string, data bson.D) error {
	_, err := client.Database("metrics").Collection(collectionName).InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}
