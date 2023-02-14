package puzzle_controller

import (
	"net/http"
	"wordle-server/puzzle"
	"wordle-server/puzzle/puzzle_model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func CreatePuzzle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var puzzleInput puzzle.PuzzleInput
		if err := ctx.BindJSON(&puzzleInput); err != nil {
			ctx.JSON(http.StatusBadRequest, puzzle.PuzzleErrorResponse{Message: err.Error()})
			return
		}

		if err := validate.Struct(&puzzleInput); err != nil {
			ctx.JSON(http.StatusBadRequest, puzzle.PuzzleErrorResponse{Message: err.Error()})
			return
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
