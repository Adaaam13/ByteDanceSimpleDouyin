package middleware

import (
	"log"
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
		// empty token
		if token == "" {
			c.JSON(http.StatusOK, Response{
				StatusCode: -1,
				StatusMsg:  "token doesn't exist",
			})
			c.Abort()
			return
		}

		log.Println("get token: ", token)

		userClaims, err := jwtService.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: -1,
				StatusMsg:  "invalid token",
			})
			c.Abort()
			return
		}
		c.Set("claims", userClaims)
		c.Next()
	}
}
