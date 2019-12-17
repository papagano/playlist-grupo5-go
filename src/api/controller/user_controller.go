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

	response, err := user.GetUser(userID, password)
	if err != nil {
		if err.Status == 0 {
			c.JSON(500, err)
		} else {
			c.JSON(err.Status, err)
		}
		return
	}

	c.JSON(http.StatusOK, response)
}
