package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/sirupsen/logrus"
	"gopkg.in/ffmt.v1"
)

type Nacos struct {
	Client      config_client.IConfigClient
	NamespaceId string
	Group       string
	DataId      string
}

func NewNacos(c chan interface{}, NamespaceId string, Group string, DataId string, IpAddr string, port uint64) *Nacos {
	sc := []constant.ServerConfig{
		{
			IpAddr: IpAddr,
			Port:   port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         NamespaceId, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}

	nacos := &Nacos{
		Client:      client,
		NamespaceId: NamespaceId,
		Group:       Group,
		DataId:      DataId,
	}
	go func() {
		nacos.Refresh(c)
	}()

	content := nacos.GetConfig()

	ffmt.Puts(content)
	return nacos
}

func (n *Nacos) Refresh(c chan interface{}) {
	_ = n.Client.ListenConfig(vo.ConfigParam{
		DataId: n.DataId,
		Group:  n.Group,
		OnChange: func(namespace, group, dataId, data string) {
			logrus.Info("Config is refreshing from nacos")
			c <- data
		},
	})
}

func (n *Nacos) GetConfig() string {
	config, err := n.Client.GetConfig(vo.ConfigParam{
		DataId: n.DataId,
		Group:  n.Group,
	})
	if err != nil {
		panic(err)
	}
	return config
}

func main() {
	c := make(chan interface{})
	n := NewNacos(c, "public", "DEFAULT_GROUP", "test1", "127.0.0.1", 8848)
	conf := n.GetConfig()

	my := MySQLConfig{}
	err := json.Unmarshal([]byte(conf), &my)
	if err != nil {
		panic(err)
	}
	ffmt.Puts(my)

	changeSettings, ok := <-c
	if ok {
		fmt.Println("changeSettings:", changeSettings)
	}
	select {}
}

type MySQLConfig struct {
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Db       string `mapstructure:"db" json:"db"`
}
