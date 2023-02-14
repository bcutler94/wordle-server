package user_controller

import (
	"fmt"
	"net/http"
	"wordle-server/user"
	"wordle-server/user/user_model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userInput user.UserInput
		if err := ctx.BindJSON(&userInput); err != nil {
			ctx.JSON(http.StatusBadRequest, user.UserErrorResponse{Message: err.Error()})
			return
		}

		if err := validate.Struct(&userInput); err != nil {
			ctx.JSON(http.StatusBadRequest, user.UserErrorResponse{Message: err.Error()})
			return
		}

		res, err := user_model.WriteUser(ctx, userInput)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, user.UserErrorResponse{Message: err.Error()})
			return
		}

		id := res.InsertedID.(primitive.ObjectID).Hex()
		fmt.Println("Successfully created User with ID ", id)
		ctx.JSON(http.StatusCreated, user.UserCreateResponse{Data: struct{ ID string }{ID: id}})
	}
}
