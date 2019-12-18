package song

import (
	"encoding/json"
	"fmt"
	"github.com/playlist-grupo5-go/src/api/domain/artist"
	"github.com/playlist-grupo5-go/src/api/domain/genere"
	"github.com/playlist-grupo5-go/src/api/utils"
	"io/ioutil"
	"net/http"
)

type Song struct {
	ID     string          `json:"id"`
	Name   string          `json:"name"`
	Genere []genere.Genere `json:"genere"`
	Artist []artist.Artist `json:"artist"`
}

type Songs []Song

func (song *Song) Get() *utils.ApiError {
	if song.ID == "" {
		return &utils.ApiError{
			Message: "ID is empty",
			Status:  http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.URL_SONG, song.ID)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	//req.Header.Set("password", password)
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

	if err := json.Unmarshal(data, &song); err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (songs *Songs) GetAll() *utils.ApiError {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", utils.URL_SONG, nil)
	//req.Header.Set("password", password)
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

	if err := json.Unmarshal(data, &songs); err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}
