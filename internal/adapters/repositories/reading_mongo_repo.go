// internal/repository/device_repo.go
package repositories

import (
	"IoTPlatform/internal/domain"
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

type ReadingRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

var _ ports.IRedingRepository = (*ReadingRepository)(nil)

func NewReadingRepository(conn string) *ReadingRepository {
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

	return &ReadingRepository{
		client:     client,
		database:   database,
		collection: collection,
	}
}

// debemos cambiar el IPORTs ya que no implementa
func (r *ReadingRepository) Create(device domain.Device) error {
	//Here your code for login in mongo database
	return nil
}

func (r *ReadingRepository) Delete(ID string) error {
	//Here your code for save in mongo database
	return nil
}
