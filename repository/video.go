package repository

import (
	"sync"
)

type Video struct {
	Id         uint   `gorm:"primaryKey;autoIncrement;not null"`
	CreateTime int64  `gorm:"not null"`
	Title      string `gorm:"not null"`
	PlayUrl    string `gorm:"not null"`
	CoverUrl   string `gorm:"not null"`
	AuthorId   uint   `gorm:"not null"`
	Visible    bool   `gorm:"not null"`
}

type VideoDao struct{}

var (
	videoDao  *VideoDao
	videoOnce sync.Once
)

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

// Create
func (*VideoDao) CreateVideo(video *Video) (*Video, error) {
	if err := db.Create(video).Error; err != nil {
		return video, err
	}
	return video, nil
}

// Retrive
func (*VideoDao) GetVideoById(video_id uint) (*Video, error) {
	var res Video
	if err := db.Where("id = ? and visible = true", video_id).First(&res).Error; err != nil {
		return &res, err
	}
	return &res, nil
}

func (*VideoDao) GetAllVideosBefore(latestTime int64) ([]*Video, error) {
	var res []*Video
	if err := db.Order("create_time desc").Where("create_time < ? and visible = true", latestTime).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (*VideoDao) GetVideosByAuthor(author_id uint) ([]*Video, error) {
	var res []*Video
	if err := db.Where("author_id = ? and visible = true", author_id).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// Update
func (*VideoDao) UpdateVideo(video_id uint, video *Video) error {
	var sourceVideo Video
	if err := db.Where("id = ?", video_id).First(&sourceVideo).Error; err != nil {
		return err
	}
	db.Model(sourceVideo).Updates(*video)
	return nil
}

// Delete
func (*VideoDao) DeleteVideo(video_id uint) error {
	if err := db.Where("video_id = ?", video_id).Delete(Video{}).Error; err != nil {
		return err
	}
	return nil
}
