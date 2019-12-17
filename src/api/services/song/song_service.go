package song

import (
	"github.com/playlist-grupo5-go/src/api/domain/song"
	"github.com/playlist-grupo5-go/src/api/utils"
)

func GetSong(songID string) (*song.Song, *utils.ApiError) {

	theSong := song.Song{
		ID: songID,
	}

	if err := theSong.Get(); err != nil {
		return nil, err
	}

	return &theSong, nil
}

func GetAllSongs() (*song.Songs, *utils.ApiError) {
	allSongs := song.Songs{}

	if err := allSongs.GetAll(); err != nil {
		return nil, err
	}

	return &allSongs, nil
}
