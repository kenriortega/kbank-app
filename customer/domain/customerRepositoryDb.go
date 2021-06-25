package customer

import (
	"context"
	"os"
	"sync"

	"github.org/kbank/errs"
	"github.org/kbank/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

//I have used below constants just to hold required database config's.
const (
	DB        = "kbank"
	CUSTOMERS = "customers"
)

type CustomerRepositoryDb struct {
	client *mongo.Client
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var filter bson.D
	if status == "" {
		//Define filter query for fetching specific document from collection
		filter = bson.D{{}} //bson.D{{}} specifies 'all documents'
	} else {
		// TODO: make filter query
	}
	var customers []Customer
	collection := d.client.Database(DB).Collection(CUSTOMERS)
	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		logger.Error(findError.Error())
		return customers, errs.NewNotFoundError("Customer Not found")
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := Customer{}
		err := cur.Decode(&t)
		if err != nil {
			logger.Error(err.Error())
			return customers, errs.NewUnexpectedError("Unexpected error on map result")
		}
		customers = append(customers, t)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(customers) == 0 {
		return customers, errs.NewNotContentError("no documents")
	}
	return customers, nil

}

//GetMongoClient - Return mongodb connection to work with
func NewCustomerRepositoryDb() CustomerRepositoryDb {
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
	return CustomerRepositoryDb{
		client: clientInstance,
	}
}
