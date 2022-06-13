package videoService

import (
	"errors"
	"simple-tiktok/repository"
	"simple-tiktok/service"
)

func Feed(latestTime int64, user_id uint) ([]*service.VideoInfo, error) {
	if latestTime < 0 {
		return nil, errors.New("无效latestTime")
	}
	if user_id == 0 {
		return nil, errors.New("无效用户id")
	}

	videos, err := repository.NewVideoDaoInstance().GetAllVideosBefore(latestTime)
	if err != nil {
		return nil, err
	}

	res, err := PackVideoInfo(videos, user_id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
