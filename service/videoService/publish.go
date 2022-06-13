package videoService

import (
	"errors"
	"mime/multipart"
	"simple-tiktok/repository"
	"simple-tiktok/service/ossService"
	"strconv"
	"time"
)

func Publish(file *multipart.FileHeader, title string, user_id uint) error {
	if user_id == 0 {
		return errors.New("无效用户id")
	}

	user, err := repository.NewUserDaoInstance().GetUserById(user_id)
	if err != nil {
		return err
	}

	video := &repository.Video{
		CreateTime: time.Now().Unix(),
		Title:      title,
		PlayUrl:    "",
		CoverUrl:   "",
		AuthorId:   user_id,
		Visible:    false,
	}
	video, err = repository.NewVideoDaoInstance().CreateVideo(video)
	if err != nil {
		return err
	}

	videos, err := repository.NewVideoDaoInstance().GetVideosByAuthor(user_id)
	if err != nil {
		return err
	}

	var createBucket bool = len(videos) == 0
	playUrl, coverUrl, err := upload(file, user.Username, video.Id, createBucket)
	if err != nil {
		_ = repository.NewVideoDaoInstance().DeleteVideo(video.Id)
		return err
	}

	video.PlayUrl = playUrl
	video.CoverUrl = coverUrl
	video.Visible = true
	err = repository.NewVideoDaoInstance().UpdateVideo(video.Id, video)
	if err != nil {
		_ = repository.NewVideoDaoInstance().DeleteVideo(video.Id)
		return err
	}
	return nil
}

func upload(file *multipart.FileHeader, username string, video_id uint, createBucket bool) (string, string, error) {
	bucketName := "simple-tiktok-" + username + "-bucket"
	objectName := "video/" + strconv.FormatUint(uint64(video_id), 10) + "-" + file.Filename

	if createBucket {
		err := ossService.Create(bucketName)
		if err != nil {
			return "", "", err
		}
	}

	playUrl, err := ossService.Upload(file, bucketName, objectName)
	if err != nil {
		return "", "", err
	}

	coverUrl := playUrl + "?x-oss-process=" + "video/snapshot,t_0,f_jpg,w_0,h_0"

	return playUrl, coverUrl, nil
}
