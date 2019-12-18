package user

import (
	"encoding/json"
	"fmt"
	"github.com/playlist-grupo5-go/src/api/utils"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (user *User) Get(password string) *utils.ApiError {
	if user.ID == "" {
		return &utils.ApiError{
			Message: "ID is empty",
			Status:  http.StatusBadRequest,
		}
	}

	if password == "" {
		return &utils.ApiError{
			Message: "Password is empty",
			Status:  http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.URL_USER, user.ID)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("password", password)
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

	if err := json.Unmarshal(data, &user); err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}
