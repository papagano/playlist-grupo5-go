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
	router.DELETE("playlists/:playlistID", controller.DeletePlaylistOnApi)
	router.GET("/user/playlists/:userID", controller.GetPlaylistsByUserFromAPI)
	router.GET("playlists", controller.GetAllPlaylistsFromAPI)

	router.POST("playlists", controller.PostPlaylistOnAPI)
	router.POST("playlists/:idPlaylist", controller.AddSongToPlaylist)

	_ = router.Run(port)

	/*if err != nil {
		logger.Error("Couldn't start the server", err)
	}*/
}
