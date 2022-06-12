package followService

import "simple-tiktok/repository"

func Follow(follower_id uint, user_id uint) error {
	_, err := repository.NewFollowDaoInstance().CreateFollow(follower_id, user_id)
	return err
}

func UnFollow(follower_id uint, user_id uint) error {
	return repository.NewFollowDaoInstance().DeleteFollow(follower_id, user_id)
}