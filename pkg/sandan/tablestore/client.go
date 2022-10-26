package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	utils "go_practice/config"
)

//var TableStoreClient

func TableStoreClient() *tablestore.TableStoreClient {
	endpoint := utils.OTS_TEST_ENDPOINT
	instanceName := utils.OTS_TEST_INSTANCENAME
	accessKeyId := utils.OTS_TEST_KEYID
	accessKeySecret := utils.OTS_TEST_SECRET
	fmt.Println(endpoint, instanceName, accessKeyId, accessKeySecret)
	client := tablestore.NewClient(endpoint, instanceName, accessKeyId, accessKeySecret)
	return client
}
