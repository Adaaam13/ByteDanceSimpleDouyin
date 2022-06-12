package followService

import (
	"simple-tiktok/repository"
	"simple-tiktok/service"
	"simple-tiktok/service/userService"
)

func QueryFollowerList(userId uint, qUserId uint) ([]*service.UserInfo, error) {
	followers, err := repository.NewFollowDaoInstance().GetFollowersByUserId(userId)
	if err != nil {
		return nil, err
	}

	var res []*service.UserInfo
	for _, follower := range followers {
		userInfo, err := userService.QueryUserInfo(follower.FollowerId, qUserId)
		if err != nil {
			return nil, err
		}
		res = append(res, userInfo)
	}
	return res, nil
}