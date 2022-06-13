package userService

import (
	"errors"
	"simple-tiktok/repository"
	"simple-tiktok/service"
)

func QueryUserInfo(userId uint, qUserId uint) (*service.UserInfo, error) {

	if userId == 0 {
		return nil, errors.New("无效用户id")
	}

	// 1. 获取用户信息
	user, err := repository.NewUserDaoInstance().GetUserById(userId)
	if err != nil {
		return nil, err
	}

	// 2. 获取关注数
	follows, err := repository.NewFollowDaoInstance().GetFollowsByFollowerId(userId)
	if err != nil {
		return nil, err
	}

	// 3. 获取粉丝数
	followers, err := repository.NewFollowDaoInstance().GetFollowersByUserId(userId)
	if err != nil {
		return nil, err
	}

	// 4. 获取是否关注
	var isFollow bool
	if qUserId == 0 { // 未登录查询用户
		isFollow = false
	} else {
		var err error
		isFollow, err = repository.NewFollowDaoInstance().IsFollow(qUserId, userId)
		if err != nil {
			return nil, err
		}
	}

	return &service.UserInfo{
		Id:            int64(user.Id),
		Name:          user.Username,
		FollowCount:   int64(len(follows)),
		FollowerCount: int64(len(followers)),
		IsFollow:      isFollow,
	}, nil
}
