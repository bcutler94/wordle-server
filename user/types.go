package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserInput struct {
	Username string `json:"username" validate:"required,alphanum,min=1"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
}

type UserErrorResponse struct {
	Message string `json:"message"`
}

type UserCreateResponse struct {
	Data struct {
		ID string
	} `json:"data"`
}
