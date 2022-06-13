package followService

import (
	"errors"
	"simple-tiktok/repository"
	"simple-tiktok/service"
	"simple-tiktok/service/userService"
)

func QueryFollowerList(user_id uint, qUser_id uint) ([]service.UserInfo, error) {
	if user_id == 0 || qUser_id == 0 {
		return nil, errors.New("无效用户id")
	}

	followers, err := repository.NewFollowDaoInstance().GetFollowersByUserId(user_id)
	if err != nil {
		return nil, err
	}

	var res []service.UserInfo
	for _, follower := range followers {
		userInfo, err := userService.QueryUserInfo(follower.FollowerId, qUser_id)
		if err != nil {
			return nil, err
		}
		res = append(res, *userInfo)
	}
	return res, nil
}
