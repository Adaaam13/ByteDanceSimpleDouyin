package followService

import (
	"errors"
	"simple-tiktok/repository"
	"simple-tiktok/service"
	"simple-tiktok/service/userService"
)

func QueryFollowList(user_id uint, qUser_id uint) ([]*service.UserInfo, error) {
	if user_id == 0 || qUser_id == 0 {
		return nil, errors.New("无效用户id")
	}

	follows, err := repository.NewFollowDaoInstance().GetFollowsByFollowerId(user_id)
	if err != nil {
		return nil, err
	}

	var res []*service.UserInfo
	for _, follow := range follows {
		userInfo, err := userService.QueryUserInfo(follow.UserId, qUser_id)
		if err != nil {
			return nil, err
		}
		res = append(res, userInfo)
	}
	return res, nil
}
