package main

import (
	"github.com/gin-gonic/gin"
	"github.com/playlist-grupo5-go/src/api/controller"
)

const (
	port = ":8082"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("users/:userID", controller.GetUserFromAPI)

	router.GET("songs/:songID", controller.GetSongFromAPI)
	router.GET("songs", controller.GetAllSongsFromAPI)

	router.GET("playlists/:playlistID", controller.GetPlaylistFromAPI)
	router.GET("playlists", controller.GetAllPlaylistsFromAPI)

	_ = router.Run(port)

	/*if err != nil {
		logger.Error("Couldn't start the server", err)
	}*/
}
