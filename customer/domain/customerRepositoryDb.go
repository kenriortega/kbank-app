package customer

import (
	"context"
	"time"

	"github.org/kbank/internal/errs"
	"github.org/kbank/internal/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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
		filter = bson.D{primitive.E{Key: "status", Value: status}}
	}
	var customers []Customer
	collection := d.client.Database(DB).Collection(CUSTOMERS)
	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		logger.Error(findError.Error())
		return customers, errs.NotFoundError
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
		return customers, errs.NoDocumentsError
	}
	return customers, nil

}

//GetIssuesByCode - Get All issues for collection
func (d CustomerRepositoryDb) FindOne(customerID primitive.ObjectID) (Customer, *errs.AppError) {
	result := Customer{}
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "_id", Value: customerID}}
	//Get MongoDB connection using

	//Create a handle to the respective collection in the database.
	collection := d.client.Database(DB).Collection(CUSTOMERS)
	//Perform FindOne operation & validate against the error.
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, errs.NotFoundError
	}
	//Return result without any error.
	return result, nil
}

//CreateOne - Insert a new document in the collection.
func (d CustomerRepositoryDb) CreateOne(customer Customer) (*mongo.InsertOneResult, *errs.AppError) {
	customer.ID = primitive.NewObjectID()
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()
	customer.Status = "inactive"
	//Create a handle to the respective collection in the database.
	collection := d.client.Database(DB).Collection(CUSTOMERS)
	//Perform InsertOne operation & validate against the error.
	result, err := collection.InsertOne(context.TODO(), customer)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.InsertOneError
	}
	//Return success without any error.
	return result, nil
}

//CreateMany - Insert multiple documents at once in the collection.
func (d CustomerRepositoryDb) CreateMany(list []Customer) (*mongo.InsertManyResult, *errs.AppError) {
	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	insertableList := make([]interface{}, len(list))
	for i, v := range list {
		insertableList[i] = v
	}

	//Create a handle to the respective collection in the database.
	collection := d.client.Database(DB).Collection(CUSTOMERS)
	//Perform InsertMany operation & validate against the error.
	result, err := collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.InsertOneError
	}
	//Return success without any error.
	return result, nil
}

//DeleteOne - Get customer by id for collection
func (d CustomerRepositoryDb) DeleteOne(customerID primitive.ObjectID) (*mongo.DeleteResult, *errs.AppError) {
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "_id", Value: customerID}}
	//Get MongoDB connection using

	//Create a handle to the respective collection in the database.
	collection := d.client.Database(DB).Collection(CUSTOMERS)

	//Perform DeleteOne operation & validate against the error.
	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		logger.Error(err.Error())
		return nil, errs.DeleteOneError
	}
	//Return success without any error.
	return result, nil
}

//DeleteAll - Get All issues for collection
func (d CustomerRepositoryDb) DeleteAll() (*mongo.DeleteResult, *errs.AppError) {
	//Define filter query for fetching specific document from collection
	selector := bson.D{{}} // bson.D{{}} specifies 'all documents'

	//Create a handle to the respective collection in the database.
	collection := d.client.Database(DB).Collection(CUSTOMERS)
	//Perform DeleteMany operation & validate against the error.
	result, err := collection.DeleteMany(context.TODO(), selector)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.DeleteOneError
	}
	//Return success without any error.
	return result, nil
}

func (d CustomerRepositoryDb) UpdateStatusByCustomerID(
	customerID primitive.ObjectID,
	status string,
	updatedAt time.Time,
) (*mongo.UpdateResult, *errs.AppError) {
	//Define filter query for fetching specific document from collection
	filter := bson.D{primitive.E{Key: "_id", Value: customerID}}

	//Define updater for to specifiy change to be updated.
	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "status", Value: status},
		primitive.E{Key: "updated_at", Value: updatedAt},
	}}}

	collection := d.client.Database(DB).Collection(CUSTOMERS)

	//Perform UpdateOne operation & validate against the error.
	result, err := collection.UpdateOne(context.TODO(), filter, updater)

	if err != nil {
		logger.Error(err.Error())
		return nil, errs.UpdateError
	}
	//Return success without any error.
	return result, nil
}

func NewCustomerRepositoryDb(clientInstance *mongo.Client) CustomerRepositoryDb {

	return CustomerRepositoryDb{
		client: clientInstance,
	}
}
