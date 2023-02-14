package puzzle_router

import (
	"wordle-server/puzzle/puzzle_controller"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	group := r.Group("/v1/puzzle")
	group.POST("", puzzle_controller.CreatePuzzle())
	// group.GET("/:id")
}
