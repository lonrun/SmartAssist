package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client *mongo.Client
	dbName string
}

func NewDatabase(dbName string, uri string) (*Database, error) {
	// Create MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Connect to MongoDB
	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	// Ping MongoDB to check connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &Database{client, dbName}, nil
}

func (d *Database) GetCollection(collectionName string) *mongo.Collection {
	return d.client.Database(d.dbName).Collection(collectionName)
}

func (d *Database) Disconnect() error {
	err := d.client.Disconnect(context.Background())
	if err != nil {
		return fmt.Errorf("error disconnecting from MongoDB: %w", err)
	}
	return nil
}
