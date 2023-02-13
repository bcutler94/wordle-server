package user

import (
	"context"
	"wordle-server/configs"

	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = configs.DBClient.Database("wordle").Collection("users")

func WriteUser(ctx context.Context, input UserInput) (*mongo.InsertOneResult, error) {
	return UserCollection.InsertOne(ctx, input)
}
