package entry_controller

import (
	"fmt"
	"net/http"
	"strings"
	"wordle-server/entry"
	"wordle-server/entry/entry_model"
	"wordle-server/puzzle/puzzle_model"
	"wordle-server/user/user_model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

func CreateEntry() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var entryInput entry.EntryInput
		if err := ctx.BindJSON(&entryInput); err != nil {
			ctx.JSON(http.StatusBadRequest, entry.EntryErrorResponse{Message: err.Error()})
			return
		}

		if err := validate.Struct(&entryInput); err != nil {
			ctx.JSON(http.StatusBadRequest, entry.EntryErrorResponse{Message: err.Error()})
			return
		}

		// Check that user exists in DB
		_, userErr := user_model.GetUser(ctx, entryInput.UserId)
		if userErr != nil {
			ctx.JSON(http.StatusBadRequest, entry.EntryErrorResponse{Message: userErr.Error()})
			return
		}

		// Check that puzzle exists in DB
		puzzleDoc, puzzleErr := puzzle_model.GetPuzzle(ctx, entryInput.PuzzleId)
		if userErr != nil {
			ctx.JSON(http.StatusBadRequest, entry.EntryErrorResponse{Message: puzzleErr.Error()})
			return
		}

		// Either create an entry if it exists or find existing entry _id
		var entryObjId primitive.ObjectID
		entryDoc, entryErr := entry_model.GetEntry(ctx, entryInput)
		fmt.Println("doc", entryDoc, entryErr)
		if entryErr == mongo.ErrNoDocuments {
			res, insertErr := entry_model.CreateEntry(ctx, entryInput)
			if insertErr != nil {
				ctx.JSON(http.StatusInternalServerError, entry.EntryErrorResponse{Message: insertErr.Error()})
				return
			}
			entryObjId = res.InsertedID.(primitive.ObjectID)
		} else {
			entryObjId = entryDoc.ID
		}

		// Get the entry to score now
		entryDoc, err := entry_model.GetEntryByID(ctx, entryObjId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, entry.EntryErrorResponse{Message: err.Error()})
			return
		}

		// If they have 5 guesses just return the entry
		if len(entryDoc.Guesses) == 5 {
			ctx.JSON(http.StatusOK, entry.EntrySuccessResponse{Data: entryDoc})
			return
		}

		// Do the scoring
		result := [5]entry.LetterResult{}
		for idx, char := range entryInput.Guess {
			if byte(char) == puzzleDoc.Word[idx] {
				result[idx] = entry.Correct
			} else if strings.ContainsRune(puzzleDoc.Word, char) {
				result[idx] = entry.Exists
			} else {
				result[idx] = entry.Incorrect
			}
		}

		// Update the entry in the DB with the results and guess
		update := bson.D{
			{"$push", bson.D{
				{"guesses", entryInput.Guess},
				{"results", result},
			}},
		}
		newEntry := entry_model.UpdateEntry(ctx, entryObjId, update)
		ctx.JSON(http.StatusOK, entry.EntrySuccessResponse{Data: newEntry})
	}
}
