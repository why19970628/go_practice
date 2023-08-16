package mapx

import "fmt"

func main6() {
	sliceMap := make(map[string][]int, 8)
	fmt.Println(sliceMap)
	v, ok := sliceMap["中国"]
	if ok {
		fmt.Println(v)
	} else {
		sliceMap["中国"] = make([]int, 8)
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 200
		sliceMap["中国"][2] = 300
		sliceMap["中国"][3] = 400

	}
	for key, value := range sliceMap {
		fmt.Println(key, value)
	}
}
