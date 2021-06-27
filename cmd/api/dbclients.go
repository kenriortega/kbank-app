package api

import (
	"context"
	"os"
	"sync"

	"github.org/kbank/internal/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

func GetMongoDbClient() *mongo.Client {
	var clientInstance *mongo.Client
	//Perform connection creation operation only once.
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URL"))
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			logger.Error(err.Error())

			panic(err)
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			logger.Error(err.Error())
			panic(err)
		}
		clientInstance = client
	})
	return clientInstance
}
