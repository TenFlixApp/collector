package data

import (
	"go.mongodb.org/mongo-driver/bson"
)

func GetMetrics(collection string) ([]bson.M, error) {
	return FindData(collection, bson.D{})
}

func PushMetrics(collection string, data bson.D) error {
	return InsertData(collection, data)
}
