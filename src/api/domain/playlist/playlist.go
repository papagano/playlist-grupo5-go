package playlist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/go-meli-toolkit/restful/rest"
	"github.com/playlist-grupo5-go/src/api/domain/song"
	"github.com/playlist-grupo5-go/src/api/utils"
	"io/ioutil"
	"net/http"
)

type Playlist struct {
	ID     string     `json:"id,omitempty"`
	Name   string     `json:"name"`
	User   string     `json:"user"`
	Avatar string     `json:"avatar"`
	Songs  song.Songs `json:"songs"`
}

type PlaylistPost struct {
	Name   string     `json:"name"`
	User   string     `json:"user"`
	Avatar string     `json:"avatar"`
	Songs  song.Songs `json:"songs"`
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

	if response.StatusCode > 299 {
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

func (playlist *Playlist) Save() *utils.ApiError {

	playlist.ID = ""

	client := &http.Client{}
	body, err := json.Marshal(&playlist)
	req, _ := http.NewRequest("POST", utils.URL_PLAYLIST, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)

	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if response.StatusCode > 299 {
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

func (playlist *Playlist) Delete() *utils.ApiError {
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", utils.URL_PLAYLIST, nil)
	q := req.URL.Query()
	q.Add("id", playlist.ID)
	req.URL.RawQuery = q.Encode()
	response, err := client.Do(req)

	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if response.StatusCode > 299 {
		data, _ := ioutil.ReadAll(response.Body)
		var errResponse utils.ApiError
		_ = json.Unmarshal(data, &errResponse)
		return &errResponse
	}

	return nil
}

func (playlist *Playlist) AddSongToPlaylist(idPlaylist string, idSong string) *utils.ApiError {

	url := fmt.Sprintf(utils.URL_PLAYLIST_ADD_SONG, idPlaylist, idSong)
	res := rest.Post(url, "")

	if res == nil || res.Response == nil {
		return &utils.ApiError{
			Message: "Response timeout",
			Status:  http.StatusInternalServerError,
		}

	}

	if res.StatusCode != 200 {
		data, _ := ioutil.ReadAll(res.Body)
		var errResponse utils.ApiError
		_ = json.Unmarshal(data, &errResponse)
		errResponse.Status = res.StatusCode
		return &errResponse
	}

	if err := json.Unmarshal(res.Bytes(), &playlist); err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (playlist *Playlist) DeleteSongToPlaylist(idPlaylist string, idSong string) *utils.ApiError {

	url := fmt.Sprintf(utils.URL_PLAYLIST_ADD_SONG, idPlaylist, idSong)
	res := rest.Delete(url)

	if res == nil || res.Response == nil {
		return &utils.ApiError{
			Message: "Response timeout",
			Status:  http.StatusInternalServerError,
		}

	}

	if res.StatusCode != 204 {
		data, _ := ioutil.ReadAll(res.Body)
		var errResponse utils.ApiError
		_ = json.Unmarshal(data, &errResponse)
		errResponse.Status = res.StatusCode
		return &errResponse
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

	if response.StatusCode > 299 {
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

func (playlists *Playlists) GetByUser(userID string) (*Playlists, *utils.ApiError) {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", utils.URL_PLAYLIST, nil)
	response, err := client.Do(req)

	if err != nil {
		return nil, &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if response.StatusCode > 299 {
		data, _ := ioutil.ReadAll(response.Body)
		var errResponse utils.ApiError
		_ = json.Unmarshal(data, &errResponse)
		return nil, &errResponse
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	var allPlaylists Playlists

	if err := json.Unmarshal(data, &allPlaylists); err != nil {
		return nil, &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	var userPlaylists Playlists

	for _, playlist := range allPlaylists {
		if playlist.User == userID {
			userPlaylists = append(userPlaylists, playlist)
		}
	}

	return &userPlaylists, nil
}
