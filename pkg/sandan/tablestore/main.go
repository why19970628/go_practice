package main

import (
	"pingguoxueyuan/pkg/sandan/tablestore/sample"
)

func main() {
	tsClient := TableStoreClient()
	//sample.CreateTableSample(tsClient, "sd_test")
	sample.PutRowSample(tsClient, "sd_test")

	sample.GetRowSample(tsClient, "sd_test")
}
