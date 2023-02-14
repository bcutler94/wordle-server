package entry_router

import (
	"wordle-server/entry/entry_controller"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	group := r.Group("/v1/entry")
	group.POST("", entry_controller.CreateEntry())
}
