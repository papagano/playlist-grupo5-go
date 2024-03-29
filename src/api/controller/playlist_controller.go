package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	playlistDomain "github.com/playlist-grupo5-go/src/api/domain/playlist"
	"github.com/playlist-grupo5-go/src/api/services/playlist"
	"github.com/playlist-grupo5-go/src/api/utils"
	"io/ioutil"
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

func PostPlaylistOnAPI(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)

	defer c.Request.Body.Close()

	var playlistBody playlistDomain.Playlist

	if err := json.Unmarshal(body, &playlistBody); err != nil {
		return
	}

	if err != nil {
		c.JSON(400, utils.ApiError{
			Message: "Body is missing",
			Status:  400,
		})
	}

	response, err2 := playlist.SavePlaylist(&playlistBody)

	if err2 != nil {
		if err2.Status == 0 {
			c.JSON(500, err2)
		} else {
			c.JSON(err2.Status, err2)
		}
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteSongOnPlaylistAPI(c *gin.Context, idPlaylist string, idSong string) {

	response, err2 := playlist.DeleteSongOnPlaylist(idPlaylist, idSong)

	if err2 == nil {
		{
			c.JSON(500, err2)
		}
		return
	}

	c.JSON(http.StatusNoContent, response)
}

func DeletePlaylistOnApi(c *gin.Context) {
	playlistID := c.Param(PARAM_PLAYLIST_ID)
	idSong := c.GetHeader("idSong")
	if idSong != "" {
		DeleteSongOnPlaylistAPI(c, playlistID, idSong)
	} else {
		err := playlist.DeletePlaylist(playlistID)
		if err != nil {
			if err.Status == 0 {
				c.JSON(500, err)
			} else {
				c.JSON(err.Status, err)
			}
			return
		}

		c.JSON(http.StatusNoContent, http.NoBody)
	}

}

func AddSongToPlaylist(c *gin.Context) {

	idPlaylist := c.Param("idPlaylist")
	idSong := c.GetHeader("idSong")

	response, err2 := playlist.AddSongToPlaylist(idPlaylist, idSong)

	if err2 != nil {
		if err2.Status == 0 {
			c.JSON(500, err2)
		} else {
			c.JSON(err2.Status, err2)
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
