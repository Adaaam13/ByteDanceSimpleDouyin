package controller

import (
	"net/http"
	"simple-tiktok/service"
	"simple-tiktok/service/commentService"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response
	CommentList []service.CommentInfo `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	// 1. 处理参数
	qUser_id := c.MustGet("qUser_id").(uint)

	videoIdStr := c.Query("video_id")
	video_id, err := strconv.Atoi(videoIdStr)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")

	// 2. 处理评论操作
	switch actionType {
	case "1":
		comment, err := commentService.Comment(uint(qUser_id), uint(video_id), commentText)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
				Comment: Comment{
					Id:         comment.Id,
					User:       User(comment.User),
					Content:    comment.Content,
					CreateDate: comment.CreateDate,
				}})
			return
		}

	case "2":
		commentIdStr := c.Query("comment_id")
		comment_id, err := strconv.Atoi(commentIdStr)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
		err = commentService.Uncomment(uint(video_id), uint(qUser_id), uint(comment_id))
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 0})
			return
		}

	}
	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
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

	// 2. 获取评论列表
	comments, err := commentService.QueryCommentList(uint(video_id), uint(qUser_id))
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 3. 响应
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: comments,
	})
}
