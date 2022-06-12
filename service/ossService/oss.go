package ossService

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"fmt"
)

const (
	endpoint        = "http://oss-cn-shenzhen.aliyuncs.com"
	accessKeyId     = "LTAI5tEcxUQyx6feEWAmtRSA"
	accessKeySecret = "LMERdR11kR3IutsMES2BBKeoa2qNlI"
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
	return nil
}

func Upload(fileName string, bucketName string, objectName string) (string, error) {
	// 获取存储空间
	bucket, err := OssClient.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	// 带进度条的上传。
	err = bucket.PutObjectFromFile(objectName, fileName, oss.Progress(&OssProgressListener{}))
	if err != nil {
		return "", err
	}

	url := endpoint + "/" + bucketName + "/" + objectName
	return url, nil
}

// 定义进度条监听器。
type OssProgressListener struct {
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		fmt.Printf("Transfer Started, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferDataEvent:
		fmt.Printf("\rTransfer Data, ConsumedBytes: %d, TotalBytes %d, %d%%.",
			event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
	case oss.TransferCompletedEvent:
		fmt.Printf("\nTransfer Completed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferFailedEvent:
		fmt.Printf("\nTransfer Failed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	default:
	}
}
