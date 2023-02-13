package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateUserController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userInput UserInput
		if err := ctx.BindJSON(&userInput); err != nil {
			ctx.JSON(http.StatusBadRequest, UserErrorResponse{Message: err.Error()})
			return
		}

		if err := validate.Struct(&userInput); err != nil {
			ctx.JSON(http.StatusBadRequest, UserErrorResponse{Message: err.Error()})
			return
		}

		res, err := WriteUser(ctx, userInput)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, UserErrorResponse{Message: err.Error()})
			return
		}

		id := res.InsertedID
		fmt.Println("Successfully created User with ID ", id)
		ctx.JSON(http.StatusCreated, UserCreateResponse{Data: struct{ ID interface{} }{ID: id}})
	}
}
