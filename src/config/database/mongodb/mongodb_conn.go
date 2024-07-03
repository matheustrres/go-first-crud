package mongodb

import (
	"context"
	"os"

	"github.com/matheustrres/go-first-crud/src/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URI      = "MONGODB_URI"
	MONGODB_USERS_DB = "MONGODB_USERS_DB"
)

func NewMongoDBConn(ctx context.Context) (*mongo.Database, error) {
	mongo_uri := os.Getenv(MONGODB_URI)
	mongo_db := os.Getenv(MONGODB_USERS_DB)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Mongodb connection established")

	return client.Database(mongo_db), nil
}
