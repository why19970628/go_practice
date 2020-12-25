package main

import "fmt"

// 案例
type TestStruct struct {
	Id  uint32
	Num uint32
}

// 当传参时，函数接收到的参数是数组切片的一个复制，虽然两个是不同的变量，但是它们都有一个指向同一个地址空间的array指针，
// 当修改一个数组切片时，另外一个也会改变，所以数组切片看起来是引用传递，其实是值传递。
func main() {
	objs := []TestStruct{TestStruct{Id: 1}, TestStruct{Id: 1}}
	fmt.Println(objs) // 输出: [{1 0} {1 0}]
	test2(objs)
	fmt.Println(objs) // 输出: [{1 3} {1 0}]
}

func test2(objs []TestStruct) {
	objmap := make(map[uint32]TestStruct)
	for _, item := range objs {
		item.Num = 3           // item是临时变量,使用了新的地址,所以此处不会直接修改外层slice元素的值
		objmap[item.Id] = item // 当将临时变量加入map集合,集合中则存有改变后的元素
	}
	fmt.Println("objmap:", objmap, len(objmap))    //objmap: map[1:{1 3}] 1
	objs = objs[:0]                   // 截取后再添加,此时与外面的slice使用的是仍是同一个内存地址
	fmt.Println("objs[:0]", objs[:0]) //[]
	for _, v := range objmap {        // objmap只有一个元素,所以只修改了第一个元素
		objs = append(objs, v) // 此时再append则修改了外层的slice
	}
	fmt.Println(objs, &objs)
}

/* 注意点：(1) append()扩容后则使用的不再是同一个内存地址
 (2) 临时变量会开辟新的内存地址,不会改变原有元素的值
 (3) 切片截取后,底层使用的内存地址并不会改变
 (4) 为防止不小心在函数内改变了slice参数的外层值，最好的做法还是拷贝参数到一个新的slice进行操作，防止误操作。
 (5) 拷贝操作参考如下:
 		newobjs := make([]TestStruct, len(objs))
		copy(newobjs, objs)
*/
