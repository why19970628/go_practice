package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// srcFile could be a single file or a directory
func Zip(srcFile string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		//header.Name = path
		//if info.IsDir() {
		//	header.Name += "/"
		//} else {
		header.Method = zip.Deflate
		//}
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}

func main() {
	str := time.Now().Format("20060102-15-04-05")
	basePath := "/tmp/"
	filePath := strings.Join([]string{basePath, str}, "")
	// 下载到这里
	fmt.Println(filePath)
	if err := os.Mkdir(filePath, os.ModePerm); err == nil {
		zipfilePath := filePath + ".zip"
		Zip(filePath, zipfilePath)
		// 上传zip
		// 移除zip
		time.Sleep(time.Second * 5)
		os.Remove(zipfilePath)
		os.Remove(filePath)
	}
}

func openFile(name string) *os.File {
	//hello.txt 中的内容为 world
	f, _ := os.Open(name)
	return f
}
