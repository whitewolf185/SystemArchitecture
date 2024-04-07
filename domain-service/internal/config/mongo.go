package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoDbName = "domain"

	RouteCollectionName     = "routes"
	CompanionCollectionName = "companions"
)

func NewMongoDbConnection(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(GetValue(DbDsn)))
	if err != nil || client == nil {
		return nil, fmt.Errorf("cannot connect to mongo DB: %w", err)
	}
	return client, nil
}
