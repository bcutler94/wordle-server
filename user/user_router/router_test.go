package user_router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"wordle-server/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	godotenv.Load()
	os.Setenv("MONGODB_DB", "wordle-test")
	r := gin.Default()
	Router(r)

	w := httptest.NewRecorder()

	userInput := user.UserInput{
		Username: "test",
	}
	jsonValue, _ := json.Marshal(userInput)
	buff := bytes.NewBuffer(jsonValue)
	req, _ := http.NewRequest("POST", "/user", buff)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	os.Unsetenv("MONGO_DB")
}
