package favService

import "simple-tiktok/repository"

func Favorite(user_id uint, video_id uint) error {
	_, err := repository.NewFavDaoInstance().CreateFav(user_id, video_id)
	if err != nil {
		return err
	}
	return nil
}

func UnFavorite(user_id uint, video_id uint) error {
	err := repository.NewFavDaoInstance().DeleteFav(user_id, video_id)
	if err != nil {
		return err
	}
	return nil
}