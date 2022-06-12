package videoService

import (
	"simple-tiktok/repository"
	"simple-tiktok/service"
)

func QueryFavList(user_id uint) ([]*service.VideoInfo, error) {
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

	res, err := PackVideoInfo(videos, user_id)
	if err != nil {
		return nil, err
	}

	return res, nil
}