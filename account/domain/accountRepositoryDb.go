package account

import (
	"context"
	"time"

	"github.org/kbank/internal/errs"
	"github.org/kbank/internal/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//I have used below constants just to hold required database config's.
const (
	DB       = "kbank"
	ACCOUNTS = "accounts"
)

type AccountRepositoryDb struct {
	client *mongo.Client
}

func NewAccountRepositoryDb(clientInstance *mongo.Client) AccountRepositoryDb {

	return AccountRepositoryDb{
		client: clientInstance,
	}
}

//CreateOne - Insert a new document in the collection.
func (a AccountRepositoryDb) CreateOne(account Account) (*mongo.InsertOneResult, *errs.AppError) {
	account.ID = primitive.NewObjectID()
	account.OpeningDate = time.Now()
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()
	account.Status = "inactive"
	//Create a handle to the respective collection in the database.
	collection := a.client.Database(DB).Collection(ACCOUNTS)
	//Perform InsertOne operation & validate against the error.
	result, err := collection.InsertOne(context.TODO(), account)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.InsertOneError
	}
	//Return success without any error.
	return result, nil
}
