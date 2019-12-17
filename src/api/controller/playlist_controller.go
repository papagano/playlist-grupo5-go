package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/playlist-grupo5-go/src/api/services/playlist"
	"net/http"
)

const (
	PARAM_PLAYLIST_ID = "playlistID"
)

func GetPlaylistFromAPI(c *gin.Context) {

	playlistID := c.Param(PARAM_PLAYLIST_ID)

	response, err := playlist.GetPlaylist(playlistID)
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

func GetAllPlaylistsFromAPI(c *gin.Context) {
	response, err := playlist.GetAllPlaylists()
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

func GetPlaylistsByUserFromAPI(c *gin.Context) {

	userID := c.Param(PARAM_USER_ID)

	response, err := playlist.GetPlaylistsByUser(userID)
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
