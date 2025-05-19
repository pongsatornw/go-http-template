package mongoutils

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

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
