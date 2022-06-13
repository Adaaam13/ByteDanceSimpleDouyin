package videoService

import (
	"simple-tiktok/repository"
	"simple-tiktok/service"
	"simple-tiktok/service/userService"
)

func PackVideoInfo(videos []*repository.Video, user_id uint) ([]service.VideoInfo, error) {
	var res []service.VideoInfo
	for _, video := range videos {
		// 1. 作者信息
		userInfo, err := userService.QueryUserInfo(video.AuthorId, user_id)
		if err != nil {
			return nil, err
		}
		// 2. 点赞数
		numFavs, err := repository.NewFavDaoInstance().GetNumFavsByVideoId(video.Id)
		if err != nil {
			return nil, err
		}
		// 3. 评论数
		numComments, err := repository.NewCommentDaoInstance().GetNumCommentsByVideoId(video.Id)
		if err != nil {
			return nil, err
		}
		// 4. 用户是否点赞
		var isFav bool
		if user_id == 0 { // 未登录用户
			isFav = false
		} else { // 登录用户
			var err error
			isFav, err = repository.NewFavDaoInstance().IsFav(user_id, video.Id)
			if err != nil {
				return nil, err
			}
		}

		res = append(res, service.VideoInfo{
			Id:            int64(video.Id),
			Author:        *userInfo,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: numFavs,
			CommentCount:  numComments,
			IsFavorite:    isFav,
		})
	}

	return res, nil
}
