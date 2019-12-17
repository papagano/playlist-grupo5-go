package playlist

import (
	"encoding/json"
	"fmt"
	"github.com/playlist-grupo5-go/src/api/domain/song"
	"github.com/playlist-grupo5-go/src/api/utils"
	"io/ioutil"
	"net/http"
)

type Playlist struct {
	ID    string     `json:"id"`
	Name  string     `json:"name"`
	User  string     `json:"user"`
	Songs song.Songs `json:"songs"`
}

type Playlists []Playlist

func (playlist *Playlist) Get() *utils.ApiError {
	if playlist.ID == "" {
		return &utils.ApiError{
			Message: "ID is empty",
			Status:  http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.URL_PLAYLIST, playlist.ID)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	response, err := client.Do(req)

	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if response.StatusCode != 200 {
		data, _ := ioutil.ReadAll(response.Body)
		var errResponse utils.ApiError
		_ = json.Unmarshal(data, &errResponse)
		return &errResponse
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &playlist); err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (playlists *Playlists) GetAll() *utils.ApiError {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", utils.URL_PLAYLIST, nil)
	response, err := client.Do(req)

	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if response.StatusCode != 200 {
		data, _ := ioutil.ReadAll(response.Body)
		var errResponse utils.ApiError
		_ = json.Unmarshal(data, &errResponse)
		return &errResponse
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &playlists); err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}