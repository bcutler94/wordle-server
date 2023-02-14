package puzzle_model

import (
	"context"
	"wordle-server/configs"
	"wordle-server/puzzle"

	"go.mongodb.org/mongo-driver/mongo"
)

var PuzzleCollection = configs.DB.Collection("puzzles")

func CreatePuzzle(ctx context.Context, puzzleInput puzzle.PuzzleInput) (*mongo.InsertOneResult, error) {
	return PuzzleCollection.InsertOne(ctx, puzzleInput)
}
