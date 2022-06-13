package controller

import (
	"net/http"
	"simple-tiktok/service/userService"
	"strconv"

	"github.com/gin-gonic/gin"
)



type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token, user_id, err := userService.UserRegisterService(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   int64(*user_id),
		Token:    *token,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token, user_id, err := userService.UserLoginService(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   int64(*user_id),
		Token:    *token,
	})
}

func UserInfo(c *gin.Context) {
	// 1. 处理参数
	userIdStr := c.Query("user_id")
	user_id, _ := strconv.Atoi(userIdStr)
	qUser_id := c.MustGet("qUser_id").(uint)

	// 2. 查找用户信息
	userInfo, err := userService.QueryUserInfo(uint(user_id), uint(qUser_id))
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     User(*userInfo),
	})
}
