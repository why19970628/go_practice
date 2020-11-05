package settings

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string

	OssAccessKey       string
	OssSecretKey       string
	OssBucket          string
	OssEndPoint        string
	OssFolderPrefix    string
	OssUrlPrefixPublic string
)

func init() {
	base_config, err := ini.Load("config/sandan.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	//LoadServer(base_config)
	LoadDb(base_config)
	//LoadQiniu(base_config)
	//LoadOss(base_config)

}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}

func LoadDb(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("admin123")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

// func LoadQiniu(file *ini.File) {
// 	AccessKey = file.Section("qiniu").Key("AccessKey").String()
// 	SecretKey = file.Section("qiniu").Key("SecretKey").String()
// 	Bucket = file.Section("qiniu").Key("Bucket").String()
// 	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()
// }

func LoadOss(file *ini.File) {
	OssAccessKey = file.Section("oss").Key("AccessKey").String()
	OssSecretKey = file.Section("oss").Key("SecretKey").String()
	OssBucket = file.Section("oss").Key("Bucket").String()
	OssEndPoint = file.Section("oss").Key("EndPoint").String()
	OssFolderPrefix = file.Section("oss").Key("FolderPrefix").String()
	OssUrlPrefixPublic = file.Section("oss").Key("UrlPrefixPublic").String()
}
