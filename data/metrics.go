package data

import (
	"go.mongodb.org/mongo-driver/bson"
)

func assertResult(res []bson.M, err error) ([]bson.M, error) {
	if res == nil {
		res = []bson.M{}
	}
	return res, err
}

func GetMetrics(collection string) ([]bson.M, error) {
	return assertResult(FindData(collection, bson.D{}))
}

func GetFilteredMetrics(collection string, filter bson.D) ([]bson.M, error) {
	return assertResult(FindData(collection, filter))
}

func GetAggregatedMetrics(collection string, pipeline bson.A) ([]bson.M, error) {
	return assertResult(AggregateData(collection, pipeline))
}

func PushMetrics(collection string, data bson.D) error {
	return InsertData(collection, data)
}
