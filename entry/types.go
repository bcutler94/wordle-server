package entry

import "go.mongodb.org/mongo-driver/bson/primitive"

type EntryInput struct {
	Guess    string `json:"guess" validate:"required,len=5,alpha"`
	PuzzleId string `json:"puzzleId" validate:"required,len=24,hexadecimal"`
	UserId   string `json:"userId" validate:"required,len=24,hexadecimal"`
}

type LetterResult string

const (
	Incorrect LetterResult = "incorrect"
	Correct   LetterResult = "correct"
	Exists    LetterResult = "exists"
)

type Entry struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	PuzzleId primitive.ObjectID `bson:"puzzleId"`
	UserId   primitive.ObjectID `bson:"userId"`
	Guesses  []string           `bson:"guesses"`
	Results  [][]LetterResult   `bson:"results"`
}

type EntryErrorResponse struct {
	Message string `json:"message"`
}

type EntrySuccessResponse struct {
	Data Entry `json:"data"`
}
