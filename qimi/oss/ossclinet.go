package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/http"
)


var (
	endpoint        = ""
	accessKeyId     = ""
	accessKeySecret = ""
	bucketName      = ""
)

func handleError(err error) {
	fmt.Println("Error:", err)
	//os.Exit(-1)
}


//实例
func ClientOss() *oss.Client {
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	return client
}

func UploadOss(objectName, imgUrl string) bool {
	// 获取存储空间。
	bucket, err := ClientOss().Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		//os.Exit(-1)
		return false
	}

	// 上传本地文件。
	//err = bucket.PutObjectFromFile(objectName, localFile)
	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return false
	}

	err = bucket.PutObject(objectName, res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		//os.Exit(-1)
		return false
	}
	return true
}

func main() {
	imgUrl := "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1604304956046&di=a36c32cf4b1671cba2f5c23c7b3ba3a7&imgtype=0&src=http%3A%2F%2Fwww.hubei.gov.cn%2Fzwgk%2Fyjgl%2F201406%2FW020140610545584489158.jpg"
	UploadOss("test/i.jpg",imgUrl)
}