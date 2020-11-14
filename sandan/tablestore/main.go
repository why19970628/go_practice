package main

import (
	"pingguoxueyuan/sandan/tablestore/sample"
)
func main()  {
	tsClient := TableStoreClient()
	sample.GetRowSample(tsClient, "sd_test")
}