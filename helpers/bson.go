package helpers

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
)

func ToBsonD(m map[string]interface{}) (bson.D, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	var bsonData bson.D
	err = bson.UnmarshalExtJSON(data, true, &bsonData)
	if err != nil {
		return nil, err
	}

	return bsonData, nil
}
