package videoService

import (
	"mime/multipart"
	"simple-tiktok/repository"
	"simple-tiktok/service/ossService"
	"strconv"
	"time"
)

func Publish(file *multipart.FileHeader, title string, user_id uint) error {
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

	playUrl, coverUrl, err := upload(file, user.Username, video.Id)
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

func upload(file *multipart.FileHeader, username string, video_id uint) (string, string, error) {
	bucketName := "simple-tiktok-" + username + "-bucket"
	objectName := "video/" + strconv.FormatUint(uint64(video_id), 10) + "-" + file.Filename
	
	playUrl, err := ossService.Upload(file, bucketName, objectName)
	if err != nil {
		return "", "", err
	}

	coverUrl := ""
	// todo

	return playUrl, coverUrl, nil
}
