package data

import (
	"context"
)

func PushMetrics(collection string, data interface{}) error {
	_, err := client.Database("metrics").Collection(collection).InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}
