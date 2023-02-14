package user_model

import (
	"context"
	"wordle-server/configs"
	"wordle-server/user"

	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = configs.DB.Collection("users")

func WriteUser(ctx context.Context, input user.UserInput) (*mongo.InsertOneResult, error) {
	return UserCollection.InsertOne(ctx, input)
}
