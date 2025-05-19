package mongoutils

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestConvertStructToBson(t *testing.T) {
	req := struct {
		Username string `bson:"username"`
	}{
		Username: "test username",
	}

	result, err := StructToBson(req, WithCreatedTime(), WithUpdatedTime())
	if err != nil {
		t.Fatalf("failed to convert data to bson. %v", err)
	}

	expected := bson.M{
		"username":     "test username",
		"created_time": result["created_time"],
		"updated_time": result["updated_time"],
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected to be %v, got %v", expected, result)
	}
}
