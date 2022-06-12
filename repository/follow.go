package repository

import (
	"sync"

	"github.com/jinzhu/gorm"
)

type Follow struct {
	FollowerId uint `gorm:"not null"`
	UserId  uint `gorm:"not null"`
}

type FollowDao struct{}

var (
	followDao  *FollowDao
	followOnce sync.Once
)

func NewFollowDaoInstance() *FollowDao {
	followOnce.Do(
		func() {
			followDao = &FollowDao{}
		})
	return followDao
}

// Create
func (*FollowDao) CreateFollow(follower_id uint, user_id uint) (*Follow, error) {
	follow := &Follow{FollowerId:follower_id, UserId: user_id}
	if err := db.Create(follow).Error; err != nil {
		return follow, err
	}
	return follow, nil
}

// Retrieve
func (*FollowDao) IsFollow(follower_id uint, user_id uint) (bool, error) {
	var follow Follow
	err := db.Where("follower_id = ? and user_id = ?", follower_id, user_id).First(&follow).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (*FollowDao) GetFollowsByFollowerId(follower_id uint) ([]*Follow, error) {
	var res []*Follow
	if err := db.Where("follower_id = ?", follower_id).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (*FollowDao) GetFollowersByUserId(user_id uint) ([]*Follow, error) {
	var res []*Follow
	if err := db.Where("user_id = ?", user_id).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// Delete
func (*FollowDao) DeleteFollow(follower_id uint, user_id uint) error {
	if err := db.Where("follower_id = ? and user_id = ?", follower_id, user_id).Delete(&Follow{}).Error; err != nil {
		return err
	}
	return nil
}
