package middleware

import (
	"errors"
	"net/http"
	"simple-tiktok/service/jwtService"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Query("token")
		userClaims, err := authToken(token)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("qUser_id", userClaims.UserId)
		c.Next()
	}
}

func authToken(token string) (*jwtService.UserClaims, error) {
	if token == "" {
		return nil, errors.New("token doesn't exist")
	}

	userClaims, err := jwtService.ParseToken(token)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	return userClaims, nil
}

func JwtAuthForm() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		title := c.PostForm("title")
		userClaims, err := authToken(token)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("qUser_id", userClaims.UserId)
		c.Set("title", title)
		c.Next()
	}
}
