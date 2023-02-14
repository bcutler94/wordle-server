package puzzle

import "go.mongodb.org/mongo-driver/bson/primitive"

type PuzzleErrorResponse struct {
	Message string `json:"message"`
}

type PuzzleInput struct {
	Word string `json:"word" validate:"required,len=5,alpha"`
}

type Puzzle struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Word string             `bson:"word"`
}

type PuzzleCreateResponse struct {
	Data struct {
		ID string
	} `json:"data"`
}
