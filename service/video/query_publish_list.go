package videoService

import (
	"simple-tiktok/repository"
	"simple-tiktok/service"
)

func QueryPublishList(user_id uint) ([]*service.VideoInfo, error) {
	videos, err := repository.NewVideoDaoInstance().GetVideosByAuthor(user_id)
	if err != nil {
		return nil, err
	}

	res, err := PackVideoInfo(videos, user_id)
	if err != nil {
		return nil, err
	}

	return res, nil
}