package playlist

import (
	"github.com/playlist-grupo5-go/src/api/domain/playlist"
	"github.com/playlist-grupo5-go/src/api/utils"
)

func GetPlaylist(id string) (*playlist.Playlist, *utils.ApiError) {

	thePlaylist := playlist.Playlist{
		ID: id,
	}

	if err := thePlaylist.Get(); err != nil {
		return nil, err
	}

	return &thePlaylist, nil
}

func SavePlaylist(body []byte) (*playlist.Playlist, *utils.ApiError) {
	newPlaylist := playlist.Playlist{}

	err := newPlaylist.Save(body)

	if err != nil {
		return nil, err
	}

	return &newPlaylist, nil
}

func GetAllPlaylists() (*playlist.Playlists, *utils.ApiError) {
	allPlaylists := playlist.Playlists{}

	if err := allPlaylists.GetAll(); err != nil {
		return nil, err
	}

	return &allPlaylists, nil
}

func GetPlaylistsByUser(userID string) (*playlist.Playlists, *utils.ApiError) {
	var allPlaylists *playlist.Playlists

	allPlaylists, err := allPlaylists.GetByUser(userID)

	if err != nil {
		return nil, err
	}

	return allPlaylists, nil
}
