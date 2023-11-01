package jsonx

import "encoding/json"

func GetMarshal(i interface{}) string {
	bs, _ := json.Marshal(i)
	return string(bs)
}
