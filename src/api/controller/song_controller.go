package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/playlist-grupo5-go/src/api/services/song"
	"net/http"
)

const (
	PARAM_SONG_ID = "songID"
)

func GetSongFromAPI(c *gin.Context) {

	userID := c.Param(PARAM_SONG_ID)

	response, err := song.GetSong(userID)
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

func GetAllSongsFromAPI(c *gin.Context) {
	response, err := song.GetAllSongs()
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
