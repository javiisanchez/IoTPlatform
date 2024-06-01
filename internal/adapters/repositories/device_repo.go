// internal/repository/device_repo.go
package repositories

import (
	"IoTPlatform/internal/ports"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	MongoClientTimeout = 5
)

type DeviceRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

var _ ports.IDeviceRepository = (*DeviceRepository)(nil)

func NewDeviceRepository(conn string) *DeviceRepository {
	ctx, cancelFunc := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancelFunc()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		conn,
	))
	if err != nil {
		return nil
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil
	}

	database := client.Database("goHexagonalBlog")
	collection := database.Collection("users")

	return &DeviceRepository{
		client:     client,
		database:   database,
		collection: collection,
	}
}

func (r *DeviceRepository) Create(ID string) error {
	//Here your code for login in mongo database
	return nil
}

func (r *DeviceRepository) Delete(ID string) error {
	//Here your code for save in mongo database
	return nil
}
