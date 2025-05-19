package mongoutils

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type StructToBsonOption func(bson.M)

func WithCreatedTime() StructToBsonOption {
	return func(bsonMap bson.M) {
		bsonMap["created_time"] = time.Now()
	}
}

func WithUpdatedTime() StructToBsonOption {
	return func(bsonMap bson.M) {
		bsonMap["updated_time"] = time.Now()
	}
}

func StructToBson[T any](input T, options ...StructToBsonOption) (bson.M, error) {
	bytes, err := bson.Marshal(input)
	if err != nil {
		return nil, err
	}

	var bMap bson.M
	err = bson.Unmarshal(bytes, &bMap)
	if err != nil {
		return nil, err
	}

	for _, opt := range options {
		opt(bMap)
	}

	return bMap, nil
}
