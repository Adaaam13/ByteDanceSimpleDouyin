package userService

import (
	"errors"
	"simple-tiktok/service/jwtService"
	"simple-tiktok/repository"

	"github.com/jinzhu/gorm"
)

func UserRegisterService(username string, password string) (*string, *uint, error) {
	if username == "" || password == "" {
		return nil, nil, errors.New("invalid username or password")
	}
	_, err := repository.NewUserDaoInstance().GetUserByUsername(username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, nil, err
	} else if err == nil {
		return nil, nil, errors.New("existed user")
	}

	user, err := repository.NewUserDaoInstance().CreateUser(username, password)
	if err != nil {
		return nil, nil, err
	}

	token, err := jwtService.GenToken(username, user.Id)
	if err != nil {
		return nil, nil, err
	}
	return &token, &user.Id, nil
}