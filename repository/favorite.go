package repository

import (
	"sync"

	"github.com/jinzhu/gorm"
)

type Favorite struct {
	UserId  uint `gorm:"primaryKey;"`
	VideoId uint `gorm:"primaryKey;"`
}

type FavDao struct{}

var (
	favDao  *FavDao
	favOnce sync.Once
)

func NewFavDaoInstance() *FavDao {
	favOnce.Do(
		func() {
			favDao = &FavDao{}
		})
	return favDao
}

// Create
func (*FavDao) CreateFav(user_id uint, video_id uint) (*Favorite, error) {
	fav := &Favorite{UserId: user_id, VideoId: video_id}
	if err := db.Create(fav).Error; err != nil {
		return fav, err
	}
	return fav, nil
}

// Retrieve
func (*FavDao) IsFav(user_id uint, video_id uint) (bool, error) {
	var fav Favorite
	err := db.Where("user_id = ? and video_id = ?", user_id, video_id).First(&fav).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (*FavDao) GetFavsByUserId(user_id uint) ([]*Favorite, error) {
	var res []*Favorite
	if err := db.Where("user_id = ?", user_id).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (*FavDao) GetNumFavsByVideoId(video_id uint) (int64, error) {
	var numFavs int64
	if err := db.Model(&Favorite{}).Where("video_id = ?", video_id).Count(&numFavs).Error; err != nil {
		return -1, err
	}
	return numFavs, nil
}

// Delete
func (*FavDao) DeleteFav(user_id uint, video_id uint) error {
	if err := db.Where("user_id = ? and video_id = ?", user_id, video_id).Delete(Favorite{}).Error; err != nil {
		return err
	}
	return nil
}
