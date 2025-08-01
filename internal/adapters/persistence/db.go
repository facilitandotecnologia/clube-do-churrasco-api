package persistence

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB(config DBConfig) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.Mongodb.URI)
	if config.Mongodb.Username != "" && config.Mongodb.Password != "" {
		clientOptions.SetAuth(options.Credential{
			Username: config.Mongodb.Username,
			Password: config.Mongodb.Password,
		})
	}

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}