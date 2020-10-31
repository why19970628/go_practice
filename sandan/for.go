package main

import (
	"fmt"
	"strings"
)

func main()  {
	dyn_ids := "1,2,3,4"
	list := strings.Split(dyn_ids, ",")
	//判断redis是否有改dyn_id动态。有则剔除该id
	for i := 0; i < len(list); i++ {
		//is_exists, _ := models.Redis.Hander.Exists(redisKey).Result()
		if int64(1) == int64(1) {

			list = append(list[:i], list[i+1:]...)
			i--
		}
	}
	fmt.Println(list)


	array := []int{1,2,3}
	slice := array[1:2]

	slice2 := array[0:2]
	fmt.Println(slice)
	fmt.Println(slice2)
	slice = append(slice, slice2...)
	fmt.Println(slice)
}
