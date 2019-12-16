package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/playlist-grupo5-go/src/api/services/user"
	"net/http"
)

const (
	PARAM_USER_ID   = "userID"
	HEADER_PASSWORD = "password"
)

func GetUserFromAPI(c *gin.Context) {

	userID := c.Param(PARAM_USER_ID)
	password := c.GetHeader(HEADER_PASSWORD)

	fmt.Println(userID + " " + password)

	response, err2 := user.GetUser(userID, password)
	if err2 != nil {
		c.JSON(err2.Status, err2)
		return
	}

	c.JSON(http.StatusOK, response)
}
