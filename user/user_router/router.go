package user_router

import (
	"wordle-server/user/user_controller"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	group := router.Group("/v1/user")
	group.POST("", user_controller.CreateUser())
	// group.GET("/:id")
}
