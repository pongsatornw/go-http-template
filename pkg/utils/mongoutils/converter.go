package mongoutils

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type StructToBsonOption func(bson.M)

// WithCreatedTime returns a StructToBsonOption that sets the "created_time" field
// in the bson.M map to the current time. This is typically used when creating a new document.
func WithCreatedTime() StructToBsonOption {
	return func(bsonMap bson.M) {
		bsonMap["created_time"] = time.Now()
	}
}

// WithUpdatedTime returns a StructToBsonOption that sets the "updated_time" field
// in the bson.M map to the current time. This is typically used when updating an existing document.

func WithUpdatedTime() StructToBsonOption {
	return func(bsonMap bson.M) {
		bsonMap["updated_time"] = time.Now()
	}
}

// StructToBson converts a generic struct `input` of any type `T` into a bson.M map,
// applying any provided StructToBsonOption functions to modify the result.
//
// It first marshals the struct into BSON bytes and then unmarshals it into a bson.M.
// The optional modifier functions can be used to inject or override fields (e.g., timestamps).
//
// Returns the resulting bson.M or an error if marshalling/unmarshalling fails.
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
