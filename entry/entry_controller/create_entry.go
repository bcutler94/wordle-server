package entry_controller

import (
	"net/http"
	"wordle-server/entry"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

		ctx.JSON(http.StatusInternalServerError, entry.EntryErrorResponse{Message: "Unimplemented"})
	}
}
