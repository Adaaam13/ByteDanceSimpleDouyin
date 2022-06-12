package commentService

import (
	"simple-tiktok/repository"
	"simple-tiktok/service"
	"simple-tiktok/service/userService"
)

func QueryCommentList(video_id uint, user_id uint) ([]*service.CommentInfo, error) {
	comments, err := repository.NewCommentDaoInstance().GetCommentsByVideoId(video_id)
	if err != nil {
		return nil, err
	}
	var res []*service.CommentInfo
	for _, comment := range comments {
		userInfo, err := userService.QueryUserInfo(comment.UserId, user_id)
		if err != nil {
			return nil, err
		}
		res = append(res, &service.CommentInfo{
			Id: int64(comment.Id),
			Content: comment.Content,
			User: *userInfo,
			CreateDate: comment.CreateDate,
		})
	}
	return res, nil
}