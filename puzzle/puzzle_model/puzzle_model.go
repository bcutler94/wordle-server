package puzzle_model

import (
	"context"
	"wordle-server/configs"
	"wordle-server/puzzle"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var PuzzleCollection = configs.DB.Collection("puzzles")

func CreatePuzzle(ctx context.Context, puzzleInput puzzle.PuzzleInput) (*mongo.InsertOneResult, error) {
	return PuzzleCollection.InsertOne(ctx, puzzleInput)
}
func GetPuzzle(ctx context.Context, id string) (puzzle.Puzzle, error) {
	var puzzle puzzle.Puzzle
	objId, objIdErr := primitive.ObjectIDFromHex(id)
	if objIdErr != nil {
		return puzzle, objIdErr
	}
	filter := bson.D{{Key: "_id", Value: objId}}
	err := PuzzleCollection.FindOne(ctx, filter).Decode(&puzzle)
	return puzzle, err
}
