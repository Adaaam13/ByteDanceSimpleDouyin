package repository

import (
	"sync"
	"time"
)

type User struct {
	Id         uint   `gorm:"primaryKey;autoIncrement;not null"`
	CreateTime int64  `gorm:"not null"`
	Username   string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
}

type UserDao struct{}

var (
	userDao  *UserDao
	userOnce sync.Once
)

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

// Create
func (*UserDao) CreateUser(username string, password string) (*User, error) {
	user := &User{CreateTime: time.Now().Unix(), Username: username, Password: password}
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Retrieve
func (*UserDao) GetUserById(userId uint) (*User, error) {
	var res User
	if err := db.Where("id = ?", userId).First(&res).Error; err != nil {
		return &res, err
	}
	return &res, nil
}

func (*UserDao) GetUserByUsername(username string) (*User, error) {
	var res User
	if err := db.Where("username = ?", username).First(&res).Error; err != nil {
		return &res, err
	}
	return &res, nil
}
