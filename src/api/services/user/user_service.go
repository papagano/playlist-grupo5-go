package user

import (
	"github.com/playlist-grupo5-go/src/api/domain/user"
	"github.com/playlist-grupo5-go/src/api/utils"
)

func GetUser(userID string, password string) (*user.User, *utils.ApiError) {

	theUser := user.User{
		ID: userID,
	}

	if err := theUser.Get(password); err != nil {
		return nil, err
	}

	return &theUser, nil
}
