package favService

import (
	"errors"
	"simple-tiktok/repository"
)

func Favorite(user_id uint, video_id uint) error {
	if user_id == 0 || video_id == 0 {
		return errors.New("无效用户id或视频id")
	}

	_, err := repository.NewFavDaoInstance().CreateFav(user_id, video_id)
	if err != nil {
		return err
	}
	return nil
}

func UnFavorite(user_id uint, video_id uint) error {
	if user_id == 0 || video_id == 0 {
		return errors.New("无效用户id或视频id")
	}

	err := repository.NewFavDaoInstance().DeleteFav(user_id, video_id)
	if err != nil {
		return err
	}
	return nil
}