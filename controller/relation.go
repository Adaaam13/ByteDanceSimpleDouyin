package controller

import (
	"net/http"
	"simple-tiktok/service"
	"simple-tiktok/service/followService"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []service.UserInfo `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	// 1. 处理参数
	qUser_id := c.MustGet("qUser_id").(uint)
	toUserIdStr := c.Query("to_user_id")
	to_user_id, _ := strconv.Atoi(toUserIdStr)

	actionType := c.Query("action_type")

	// 2. 处理关注操作
	switch actionType {
	case "1":
		err := followService.Follow(uint(qUser_id), uint(to_user_id))
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
	case "2":
		err := followService.UnFollow(uint(qUser_id), uint(to_user_id))
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

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	// 1. 处理参数
	qUser_id := c.MustGet("qUser_id").(uint)
	userIdStr := c.Query("user_id")
	user_id, _ := strconv.Atoi(userIdStr)

	// 2. 关注列表
	userInfos, err := followService.QueryFollowList(uint(user_id), uint(qUser_id))
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 3. 响应
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userInfos,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	// 1. 处理参数
	qUser_id := c.MustGet("qUser_id").(uint)
	userIdStr := c.Query("user_id")
	user_id, _ := strconv.Atoi(userIdStr)

	// 2. 粉丝列表
	userInfos, err := followService.QueryFollowerList(uint(user_id), uint(qUser_id))
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 3. 响应
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userInfos,
	})
}
