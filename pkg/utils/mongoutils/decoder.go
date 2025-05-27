package mongoutils

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// DecodeFromCursor reads all documents from the given MongoDB cursor and decodes
// them into a slice of the specified generic type `Output`.
//
// It automatically closes the cursor after reading. If any decoding error occurs,
// the function returns nil and the error.
//
// Example usage:
//
//	var results []MyStruct
//	results, err := DecodeFromCursor[MyStruct](cursor)
//
// Returns a slice of decoded documents or an error.
func DecodeFromCursor[Output any](cursor *mongo.Cursor) ([]Output, error) {
	defer func() { _ = cursor.Close(context.Background()) }()

	outputs := make([]Output, 0, cursor.RemainingBatchLength())
	for cursor.Next(context.Background()) {
		var out Output
		if err := cursor.Decode(&out); err != nil {
			return nil, err
		}

		outputs = append(outputs, out)
	}

	return outputs, nil
}
