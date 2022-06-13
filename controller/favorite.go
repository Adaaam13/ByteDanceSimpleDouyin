package controller

import (
	"net/http"
	favService "simple-tiktok/service/favoriteService"
	"simple-tiktok/service/videoService"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	// 1. 处理参数
	qUser_id := c.MustGet("qUser_id").(uint)

	videoIdStr := c.Query("video_id")
	video_id, err := strconv.ParseUint(videoIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	actionType := c.Query("action_type")

	// 2. 处理点赞操作
	switch actionType {
	case "1":
		err := favService.Favorite(uint(qUser_id), uint(video_id))
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
	case "2":
		err := favService.UnFavorite(uint(qUser_id), uint(video_id))
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
	}
	// 3. 响应
	c.JSON(http.StatusOK, Response{StatusCode: 0})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	// 1. 处理参数
	qUser_id := c.MustGet("qUser_id").(uint)

	userIdStr := c.Query("user_id")
	user_id, _ := strconv.Atoi(userIdStr)

	// 2. 获取用户点赞列表
	videos, err := videoService.QueryFavList(uint(user_id), uint(qUser_id))
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
