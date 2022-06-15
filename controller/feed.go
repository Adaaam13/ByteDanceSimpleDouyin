package controller

import (
	"net/http"
	"simple-tiktok/service"
	"simple-tiktok/service/jwtService"
	"simple-tiktok/service/videoService"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []service.VideoInfo `json:"video_list,omitempty"`
	NextTime  int64               `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// 1. 处理参数
	// 1.1 latestTime(string) -> latest_time(int64)
	latestTime := c.Query("latest_time")
	var latest_time int64
	if latestTime == "" {
		latest_time = time.Now().Unix()
	} else {
		var err error
		latest_time, err = strconv.ParseInt(latestTime, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response:  Response{StatusCode: -1},
				VideoList: nil,
				NextTime:  time.Now().Unix(),
			})
		}
	}
	// 1.2 token(string) -> user_id(uint)
	token := c.Query("token")
	var user_id uint
	if token == "" {
		user_id = 0
	} else {
		claims, err := jwtService.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response:  Response{StatusCode: -1},
				VideoList: nil,
				NextTime:  time.Now().Unix(),
			})
		}
		user_id = claims.UserId
	}

	// 2. 获取视频流
	videos, err := videoService.Feed(latest_time, user_id)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: -1},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
	}

	// 3. 响应
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
