package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/ffmt.v1"
	"log"
	"time"
)

type NacOsServerConfig struct {
	Scheme      string `json:"scheme"`       //the nacos server scheme
	ContextPath string `json:"context_path"` //the nacos server contextpath
	IpAddr      string `json:"ip_addr"`      //the nacos server address
	Port        uint64 `json:"port"`         //the nacos server port
}

//GetConfigJsonToEntity 配置文件是json的话直接读取成实体
func (n *NacOsServerConfig) GetConfigJsonToEntity(dataId, group string, resp interface{}) error {
	c, err := n.GetConfig(dataId, group)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(c), resp)
}

func NewNacOsServerConfig(IpAddr, ContextPath string, Port uint64) *NacOsServerConfig {
	return &NacOsServerConfig{
		IpAddr:      IpAddr,
		ContextPath: ContextPath,
		Port:        Port,
	}
}

// GetConfig 读取string
func (n *NacOsServerConfig) GetConfig(dataId, group string) (string, error) {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      n.IpAddr,
			ContextPath: n.ContextPath,
			Port:        n.Port,
			Scheme:      n.Scheme,
		},
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 监听配置
	configClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("ListenConfig group:"+namespace, group+", dataId:"+dataId+", data:"+data)
		},
	})

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})

	return content, nil
}

var serverConfig *ServerConfig

//ConfigByNacOS 通过nacos做配置中心获取config
//github.com/nacos-group/nacos-sdk-go
func main() {
	t := NewNacOsServerConfig("127.0.0.1", "/nacos", 8848)

	for {
		str, err := t.GetConfig("test1", "DEFAULT_GROUP")
		if err != nil {
			panic(err)
		}

		fmt.Println(str)
		err = json.Unmarshal([]byte(str), &serverConfig)
		ffmt.Puts(serverConfig)
		time.Sleep(time.Second * 10)
	}

}

type RabbitConfig struct {
	Host  string `mapstructure:"host" json:"host"`
	Port  int    `mapstructure:"port" json:"port"`
	Vhost string `mapstructure:"vhost" json:"vhost"`
}

type MySQLConfig struct {
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Db       string `mapstructure:"db" json:"db"`
}

type RedisConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

type NacosInit struct {
	Nacos NacosConfig `mapstructure: "nacos"`
}

type ServerConfig struct {
	MySQL               MySQLConfig  `mapstructure:"mysql" json:"mysql"`
	Redis               RedisConfig  `mapstructure:"redis" json:"redis"`
	Rabbit              RabbitConfig `mapstructure:"rabbit" json:"rabbit"`
	UploadServicePort   int          `mapstructure:"uploadservice" json:"uploadservice"`
	DownloadServicePort int          `mapstructure:"downloadservice" json:"downloadservice"`
}
