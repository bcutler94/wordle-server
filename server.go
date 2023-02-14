package main

import (
	"wordle-server/puzzle/puzzle_router"
	"wordle-server/user/user_router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	user_router.Router(r)
	puzzle_router.Router(r)

	r.Run()
}
