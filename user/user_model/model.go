package user_model

import (
	"context"
	"wordle-server/configs"
	"wordle-server/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = configs.DB.Collection("users")

func WriteUser(ctx context.Context, input user.UserInput) (*mongo.InsertOneResult, error) {
	return UserCollection.InsertOne(ctx, input)
}

func GetUser(ctx context.Context, id string) (user.User, error) {
	var user user.User
	objId, objIdErr := primitive.ObjectIDFromHex(id)
	if objIdErr != nil {
		return user, objIdErr
	}
	filter := bson.D{{Key: "_id", Value: objId}}
	err := UserCollection.FindOne(ctx, filter).Decode(&user)
	return user, err
}
