package main

import "fmt"

// 案例
type TestStruct struct {
	Id  uint32
	Num uint32
}

// 当传参时，函数接收到的参数是数组切片的一个复制，虽然两个是不同的变量，但是它们都有一个指向同一个地址空间的array指针，
// 当修改一个数组切片时，另外一个也会改变，所以数组切片看起来是引用传递，其实是值传递。
//func main() {
//	objs := []TestStruct{TestStruct{Id: 1}, TestStruct{Id: 1}}
//	fmt.Println(objs) // 输出: [{1 0} {1 0}]
//	test2(objs)
//	fmt.Println(objs) // 输出: [{1 3} {1 0}]
//}

func test2(objs []TestStruct) {
	objmap := make(map[uint32]TestStruct)
	for _, item := range objs {
		item.Num = 3           // item是临时变量,使用了新的地址,所以此处不会直接修改外层slice元素的值
		objmap[item.Id] = item // 当将临时变量加入map集合,集合中则存有改变后的元素
	}
	fmt.Println("objmap:", objmap, len(objmap)) //objmap: map[1:{1 3}] 1
	objs = objs[:0]                             // 截取后再添加,此时与外面的slice使用的是仍是同一个内存地址
	fmt.Println("objs[:0]", objs[:0])           //[]
	for _, v := range objmap {                  // objmap只有一个元素,所以只修改了第一个元素
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

func f1(a []int) {
	a[0] = 100
	a = append(a, 30)
	fmt.Println("2:", a)
}

/*
解释：

首先，打印出 a1: [4 5 6]，这是 fn1 函数中的第一个打印语句，显示了切片 a1 的值为 [4 5 6]，这个值是从 nums 中索引为3（包括）到末尾的元素构成的。

然后，打印出 a2: [4 5 7]，这是 fn1 函数中的第二个打印语句，显示了修改后的切片 a1 的值为 [4 5 7]，其中索引为2的元素被修改为了7。

最后，打印出 fn1 after nums: [1 2 3 4 5 7]，这是 sliceQuestion 函数中的打印语句
，显示了函数调用后原始的 nums 切片被修改了，即索引为5的位置上的元素由原来的6被修改为了7。
这是因为在Go语言中，切片是引用类型，传递给函数的是其底层数组的引用，因此在函数内部修改切片会影响到原始切片。

*/

func sliceQuestion() {
	nums1 := []int{1, 2, 3, 4, 5, 6}
	fn1(nums1)
	fmt.Println("fn1 after nums: ", nums1) // 1 2 3 4 5 7

	nums2 := []int{1, 2, 3, 4, 5, 6}
	fn2(nums2)
	fmt.Println("fn2 after nums: ", nums2) // 1 2 3 4 5 7

	nums3 := []int{1, 2, 3, 4, 5, 6}
	fn_append(nums3)
	fmt.Println("fn_append after nums: ", nums3) // 1 2 3 4 5 6
}
func fn1(nums []int) {
	a1 := nums[3:] // 4 5 6
	fmt.Println("a1:", a1)
	a1[2] = 7 // 4 5 7
	fmt.Println("a2:", a1)
}

func fn2(nums []int) {
	nums = nums[3:] // 4 5 6
	fmt.Println("a3:", nums)
	nums[2] = 7 // 4 5 7
	fmt.Println("a4:", nums)
}

func fn_append(nums []int) {
	nums = nums[3:]
	fmt.Printf("a5: %vlen:%v,cap:%v \n", nums, len(nums), cap(nums)) // [4 5 6] len:3,cap:3
	nums[2] = 7                                                      // 4 5 7
	nums = append(nums, 7)                                           // [4 5 7 7]，发生扩容
}

func main() {
	sliceQuestion()
	//a := make([]int, 2)
	//a[0] = 10
	//a[1] = 20
	//fmt.Println("1:", a)
	//f1(a)
	//fmt.Println(a)
}
