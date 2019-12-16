package main

import (
	"github.com/gin-gonic/gin"
	user "github.com/playlist-grupo5-go/src/api/controller"
)

const (
	port = ":8082"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("users/:userID", user.GetUserFromAPI)

	router.Run(port)
}
