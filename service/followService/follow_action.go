package followService

import (
	"errors"
	"simple-tiktok/repository"
)

func Follow(follower_id uint, user_id uint) error {
	if follower_id == 0 || user_id == 0 {
		return errors.New("无效用户id")
	}
	if follower_id == user_id {
		return errors.New("无效关注")
	}
	_, err := repository.NewFollowDaoInstance().CreateFollow(follower_id, user_id)
	return err
}

func UnFollow(follower_id uint, user_id uint) error {
	if follower_id == 0 || user_id == 0 {
		return errors.New("无效用户id")
	}
	if follower_id == user_id {
		return errors.New("无效取消关注")
	}
	return repository.NewFollowDaoInstance().DeleteFollow(follower_id, user_id)
}