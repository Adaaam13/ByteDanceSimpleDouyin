package videoService

import (
	"errors"
	"simple-tiktok/repository"
	"simple-tiktok/service"
)

func QueryFavList(user_id uint, qUser_id uint) ([]service.VideoInfo, error) {
	if user_id == 0 {
		return nil, errors.New("无效用户id")
	}

	favs, err := repository.NewFavDaoInstance().GetFavsByUserId(user_id)
	if err != nil {
		return nil, err
	}

	var videos []*repository.Video
	for _, fav := range favs {
		video, err := repository.NewVideoDaoInstance().GetVideoById(fav.VideoId)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	res, err := PackVideoInfo(videos, qUser_id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
