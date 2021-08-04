package main

import (
	"encoding/base64"
	"fmt"
	"github.com/disintegration/imageorient"
	"image"
	"io/ioutil"
	"net/http"
)

const oss_thumbnail_suffix = "?x-oss-process=style/thumbnail"

func GetImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	//ffmt.Puts(resp)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	img, _, err := imageorient.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func main() {
	image_url := "http://cdn-dev.sandanapp.com/test/6f786b3114c6677f_500_500.png"
	url := image_url + oss_thumbnail_suffix
	fmt.Println(url)

	imgUrl := "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1604304956046&di=a36c32cf4b1671cba2f5c23c7b3ba3a7&imgtype=0&src=http%3A%2F%2Fwww.hubei.gov.cn%2Fzwgk%2Fyjgl%2F201406%2FW020140610545584489158.jpg"
	//imgData, _ := ioutil.ReadFile(url2)
	//fmt.Println(imgData)
	//buf := bytes.NewBuffer(imgData)
	//image, err := imaging.Decode(buf)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//ffmt.Puts(image)
	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	imageBase64 := base64.StdEncoding.EncodeToString(data)
	fmt.Println("base64", imageBase64)

	//im, _ := GetImage(url)
	//fmt.Println(im)
	//fmt.Println(im.Bounds().Dx(), im.Bounds().Dy()) // 是我需要的宽和高

	fmt.Println("------------------")
	//url2 := "http://cdn-dev.sandanapp.com/test/46450376.png"
	//GetImage(url2)

}
