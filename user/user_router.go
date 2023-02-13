package user

import "github.com/gin-gonic/gin"

func UserRouter(router *gin.Engine) {
	group := router.Group("/v1/user")
	group.POST("", CreateUserController())
	group.GET("/:id")
}
