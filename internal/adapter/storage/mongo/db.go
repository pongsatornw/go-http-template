package mongo

import (
	"context"
	"pongsatorn/go-http-template/internal/adapter/config"
	"pongsatorn/go-http-template/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewMongoDatabse(config config.MongoConfig) (*mongo.Database, func(), error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFn()

	option := options.Client().ApplyURI(config.URI).SetConnectTimeout(config.ConnectionTimeout)
	client, err := mongo.Connect(option)
	if err != nil {
		return nil, cancelFn, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, cancelFn, err
	}

	cleanUp := func() {
		logger.Logger.Info("prepare to disconnect client from MongoDB")
		_ = client.Disconnect(context.Background())
	}

	return client.Database(config.DatabaseName), cleanUp, nil
}
