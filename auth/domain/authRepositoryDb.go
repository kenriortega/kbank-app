package auth

import (
	"context"
	"fmt"
	"time"

	"github.org/kbank/internal/cryptopass"
	"github.org/kbank/internal/errs"
	"github.org/kbank/internal/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//I have used below constants just to hold required database config's.
const (
	DB    = "kbank"
	USERS = "users"
)

type AuthRepositoryDb struct {
	client *mongo.Client
}

func NewAuthRepositoryDb(clientInstance *mongo.Client) AuthRepositoryDb {

	return AuthRepositoryDb{
		client: clientInstance,
	}
}

//CreateOne - Insert a new document in the collection.
func (a AuthRepositoryDb) CreateOne(newUser User) (*mongo.InsertOneResult, *errs.AppError) {
	newUser.ID = primitive.NewObjectID()

	newUser.Password = cryptopass.HashAndSalt([]byte(newUser.Password))
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()
	fmt.Println(newUser.Password)
	//Create a handle to the respective collection in the database.
	collection := a.client.Database(DB).Collection(USERS)
	//Perform InsertOne operation & validate against the error.
	result, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.InsertOneError
	}
	//Return success without any error.
	return result, nil
}

//CreateOne - Insert a new document in the collection.
func (a AuthRepositoryDb) Login(username, password string) (User, *errs.AppError) {

	result := User{}
	//Define filter query for fetching specific document from collection
	filter := bson.D{
		primitive.E{Key: "username", Value: username},
	}

	//Create a handle to the respective collection in the database.
	collection := a.client.Database(DB).Collection(USERS)
	//Perform FindOne operation & validate against the error.
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, errs.NotFoundError
	}
	// comapre password
	hashedPassword := result.Password
	if cryptopass.ComparePasswords(hashedPassword, []byte(password)) {
		//Return result without any error.
		return result, nil
	} else {
		//Return result without any error.
		return result, errs.NewUnexpectedError("Password no match")
	}
}
