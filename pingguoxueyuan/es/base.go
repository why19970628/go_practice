package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)


var EsClient *elastic.Client

//var host = "http://127.0.0.1:9200/"

var esUrl = "http://localhost:9200"


func init(){
	//es 配置
	var err error
	//client, err := elastic.NewClient(elastic.SetURL("Elastic_Host"))
	//if err != nil {
	//	fmt.Println("Error when connecting Elasticsearch host");
	//}
	EsClient, err := elastic.NewSimpleClient( elastic.SetURL(esUrl))
	//EsClient, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(esUrl))
	if err != nil {
		panic(err)
	}
	info, code, err := EsClient.Ping(esUrl).Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := EsClient.ElasticsearchVersion(esUrl)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
}
