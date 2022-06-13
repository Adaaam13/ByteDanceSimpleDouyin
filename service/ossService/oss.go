package ossService

import (
	"mime/multipart"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
	endpoint        = "http://oss-cn-shenzhen.aliyuncs.com"
	accessKeyId     = ""
	accessKeySecret = ""
	endpointPostFix = ".oss-cn-shenzhen.aliyuncs.com"
)

var OssClient *oss.Client

func Init() {
	var err error
	OssClient, err = oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
}

func Create(bucketName string) error {
	err := OssClient.CreateBucket(bucketName)
	if err != nil {
		return err
	}
	// default public read
	err = OssClient.SetBucketACL(bucketName, oss.ACLPublicRead)
    if err != nil {
		return err
    }
	return nil
}

func Upload(file *multipart.FileHeader, bucketName string, objectName string) (string, error) {
	// 获取存储空间
	bucket, err := OssClient.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// 带进度条的上传。
	err = bucket.PutObject(objectName, src)
	if err != nil {
		return "", err
	}

	url := "https://" + bucketName + endpointPostFix + "/" + objectName
	return url, nil
}
