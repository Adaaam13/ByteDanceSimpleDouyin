package videoService

import (
	"errors"
	"simple-tiktok/repository"
	"simple-tiktok/service"
)

func QueryPublishList(user_id uint, qUser_id uint) ([]service.VideoInfo, error) {
	if user_id == 0 {
		return nil, errors.New("无效用户id")
	}

	videos, err := repository.NewVideoDaoInstance().GetVideosByAuthor(user_id)
	if err != nil {
		return nil, err
	}

	res, err := PackVideoInfo(videos, qUser_id)
	if err != nil {
		return nil, err
	}

	return res, nil
}