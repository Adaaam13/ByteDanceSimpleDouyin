package repository

import (
	"sync"
	"time"
)

type Comment struct {
	Id         uint   `gorm:"not null;autoIncrement"`
	Content    string `gorm:"not null"`
	UserId     uint   `gorm:"not null"`
	VideoId    uint   `gorm:"not null"`
	CreateDate string  `gorm:"not null"`
}

type CommentDao struct{}

var (
	commentDao  *CommentDao
	commentOnce sync.Once
)

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}

// Create
func (*CommentDao) CreateComment(user_id uint, video_id uint, content string) (*Comment, error) {
	comment := &Comment{Content: content, UserId: user_id, VideoId: video_id, CreateDate: time.Now().Format("2006-01-02 15:04:05")[5:10]}
	if err := db.Create(comment).Error; err != nil {
		return comment, err
	}
	return comment, nil
}

// Retrieve
func (*CommentDao) GetCommentById(comment_id uint) (*Comment, error) {
	var c Comment
	if err := db.Where("id = ?", comment_id).Find(&c).Error; err != nil {
		return &c, err
	}
	return &c, nil
}

func (*CommentDao) GetCommentsByVideoId(video_id uint) ([]*Comment, error) {
	var res []*Comment
	if err := db.Where("video_id = ?", video_id).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (*CommentDao) GetNumCommentsByVideoId(video_id uint) (int64, error) {
	var numComments int64
	if err := db.Model(&Comment{}).Where("video_id = ?", video_id).Count(&numComments).Error; err != nil {
		return -1, err
	}
	return numComments, nil
}

// Delete
func (*CommentDao) DeleteCommentById(comment_id uint) error {
	if err := db.Where("id = ?", comment_id).Delete(Comment{}).Error; err != nil {
		return err
	}
	return nil
}
