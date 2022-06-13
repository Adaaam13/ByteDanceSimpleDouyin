package controller

import (
	"simple-tiktok/service/ossService"
	"simple-tiktok/repository"
)

func Init() {
	repository.Init()
	ossService.Init()
}