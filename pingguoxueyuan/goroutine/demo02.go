package main
import (
	"encoding/json"
	"golang.org/x/sync/singleflight"
	"gopkg.in/ffmt.v1"
)

var batchDyn  singleflight.Group

type Data struct {
	Name string
}
func test(){
	//从数据库获取数据
	sign := make(chan struct{})
	fn := func () (interface{}, error) {
	//从数据库获取数据 确保 single
	data := Data{Name: "wang",}
	return data, nil
	}

	new := Data{}
	go func () {
		result, _, _ := batchDyn.Do("dynBatch", fn)
		resBytes, _ := json.Marshal(result)
		_ = json.Unmarshal(resBytes, &new)
		sign <- struct{}{}
	}()
	<-sign
	ffmt.Puts(new)
}

func main() {
	test()
}