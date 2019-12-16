package user

import (
	"github.com/apigotest/src/apigo/utils"
	"github.com/playlist-grupo5-go/src/api/domain/user"
)

func GetUser(userID string, password string) (*user.User, *utils.ApiError) {

	user := user.User{
		ID: userID,
	}

	if err := user.Get(password); err != nil {
		return nil, err
	}

	return &user, nil
}
