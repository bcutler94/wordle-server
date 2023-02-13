package main

import (
	"wordle-server/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	user.UserRouter(r)

	r.Run()
}
