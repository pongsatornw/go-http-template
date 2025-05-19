package repository

import (
	"context"
	"pongsatorn/go-http-template/pkg/utils/mongoutils"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

const templateCollectionName = "template"

type TemplateRepository struct {
	collection *mongo.Collection
}

func NeTemplateRepository(db *mongo.Database) *TemplateRepository {
	return &TemplateRepository{
		collection: db.Collection(templateCollectionName),
	}
}

func (r *TemplateRepository) Temp(ctx context.Context, input any) (any, error) {
	bMap, err := mongoutils.StructToBson(0, mongoutils.WithCreatedTime())
	if err != nil {
		return nil, err
	}

	return r.collection.InsertOne(ctx, bMap)
}
