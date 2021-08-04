package main

import "fmt"

// 声明后端元素的数据结构
type volume struct {
	name  string
	quota int
}

// 声明数据类型存放元素的Slice
type volumes []volume

// 声明并赋值索引下标
var i = -1

// 轮询调度算法
func (v volumes) Select() string {
	i = (i + 1) % len(v)
	return v[i].name
}

func main() {
	v1 := volume{
		name:  "a",
		quota: 4,
	}
	v2 := volume{
		name:  "b",
		quota: 4,
	}
	v3 := volume{
		name:  "c",
		quota: 8,
	}
	vo := volumes{
		v1, v2, v3,
	}

	for j := 0; j < 20; j++ {
		fmt.Println(j+1, vo.Select())
	}
}
