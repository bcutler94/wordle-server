package puzzle_controller

import (
	"net/http"
	"wordle-server/puzzle"
	"wordle-server/puzzle/puzzle_model"
	wordgenerator "wordle-server/word_generator"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePuzzle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		randomWord, err := wordgenerator.GenerateRandom()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, puzzle.PuzzleErrorResponse{Message: err.Error()})
			return
		}
		puzzleInput := puzzle.PuzzleInput{
			Word: randomWord,
		}

		res, err := puzzle_model.CreatePuzzle(ctx, puzzleInput)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, puzzle.PuzzleErrorResponse{Message: err.Error()})
			return
		}

		id := res.InsertedID.(primitive.ObjectID).Hex()
		ctx.JSON(http.StatusCreated, puzzle.PuzzleCreateResponse{Data: struct{ ID string }{ID: id}})
	}
}
