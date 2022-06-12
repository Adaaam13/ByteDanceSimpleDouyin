package commentService

import (
	"errors"
	"simple-tiktok/repository"
)

func Comment(user_id uint, video_id uint, content string) error {
	if _, err := repository.NewCommentDaoInstance().CreateComment(user_id, video_id, content); err != nil {
		return err
	}
	return nil
}

func Uncomment(video_id uint, user_id uint, comment_id uint) error {
	comment, err := repository.NewCommentDaoInstance().GetCommentById(comment_id)
	if err != nil {
		return err
	}
	if comment.UserId != user_id || comment.VideoId != video_id {
		return errors.New("用户无权限")
	}
	if err := repository.NewCommentDaoInstance().DeleteCommentById(comment_id); err != nil {
		return err
	}
	return nil
}