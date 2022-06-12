package userService

import (
	"errors"
	"simple-tiktok/repository"
	"simple-tiktok/service/jwtService"
)

func UserLoginService(username string, password string) (*string, *uint, error) {
	if username == "" || password == "" {
		return nil, nil, errors.New("invalid username or password")
	}
	user, err := repository.NewUserDaoInstance().GetUserByUsername(username)
	if err != nil {
		return nil, nil, err
	}
	if password != user.Password {
		return nil, nil, errors.New("wrong password")
	}

	token, err := jwtService.GenToken(username, user.Id)
	if err != nil {
		return nil, nil, err
	}
	return &token, &user.Id, nil
}

