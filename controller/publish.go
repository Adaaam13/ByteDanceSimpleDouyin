package controller

import (
	"net/http"
	"simple-tiktok/service"
	"simple-tiktok/service/videoService"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []service.VideoInfo `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	// 1. 处理参数
	qUser_id := c.MustGet("qUser_id").(uint)

	title := c.Query("title")

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 2. publish服务
	if err := videoService.Publish(data, title, uint(qUser_id)); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 3. 响应
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	// 1. 处理参数
	userIdStr := c.Query("user_id")
	user_id, _ := strconv.Atoi(userIdStr)
	qUser_id := c.MustGet("qUser_id").(uint)

	// 2. publishList服务
	videos, err := videoService.QueryPublishList(uint(user_id), uint(qUser_id))
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 3. 响应
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
