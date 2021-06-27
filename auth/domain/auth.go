package auth

import (
	"time"

	"github.org/kbank/internal/errs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	Role      string             `bson:"role"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type AuthRepository interface {
	Login(username, password string) (User, *errs.AppError)
	CreateOne(User) (*mongo.InsertOneResult, *errs.AppError)
}
