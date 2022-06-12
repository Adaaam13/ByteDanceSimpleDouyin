package userService

import (
	"errors"
	"simple-tiktok/repository"
	"simple-tiktok/service"
	"sync"
)

func QueryUserInfo(userId uint, qUserId uint) (*service.UserInfo, error){
	return NewQueryUserInfo(userId, qUserId).Do()
}

func NewQueryUserInfo(userId uint, qUserId uint) *QueryUserInfoFlow {
	return &QueryUserInfoFlow{
		userId: userId,
		queryUserId: qUserId,
	}
}

type QueryUserInfoFlow struct {
	userId uint
	queryUserId uint

	userInfo *service.UserInfo
}

func (f *QueryUserInfoFlow) Do() (*service.UserInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareAndPackUserInfo(); err != nil {
		return nil, err
	}
	return f.userInfo, nil
}

func (f *QueryUserInfoFlow) checkParam() error {
	if f.userId <= 0 || f.queryUserId <= 0{
		return errors.New("user id must be greater than 0")
	}
	return nil
}
func (f *QueryUserInfoFlow) prepareAndPackUserInfo() error {
	var wg sync.WaitGroup
	wg.Add(4)

	// 1. 获取user信息
	go func() {
		defer wg.Done()
		user, _ := repository.NewUserDaoInstance().GetUserById(f.userId)
		f.userInfo.Id = int64(f.userId)
		f.userInfo.Name = user.Username
	}()

	// 2. 获取关注数
	go func() {
		defer wg.Done()
		follows, _ := repository.NewFollowDaoInstance().GetFollowsByFollowerId(f.userId)
		f.userInfo.FollowCount = int64(len(follows))
	}()
	
	// 3. 获取粉丝数
	go func() {
		defer wg.Done()
		followers, _ := repository.NewFollowDaoInstance().GetFollowersByUserId(f.userId)
		f.userInfo.FollowerCount = int64(len(followers))
	}()

	// 4. 获取是否关注
	go func() {
		defer wg.Done()
		isFollow, _ := repository.NewFollowDaoInstance().IsFollow(f.queryUserId, f.userId)
		f.userInfo.IsFollow = isFollow
	}()

	wg.Wait()
	return nil
}