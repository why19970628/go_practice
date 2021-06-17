package main

//代码
import (
	"fmt"
)

//二分查找函数 //假设有序数组的顺序是从小到大（很关键，决定左右方向）
func BinaryFind(arr *[]int, leftIndex int, rightIndex int, findVal int) {
	//判断leftIndex是否大于rightIndex
	if leftIndex > rightIndex {
		fmt.Println("没找到")
		return
	}

	//先找到中间的下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下标是%v\n", middle)
	}
}

func BinarySearch(arr []int, left int, right int, result int) {
	if left >= right {
		fmt.Println("not find")
		return
	}
	mid := (left + right) / 2
	if arr[mid] > result {
		BinarySearch(arr, left, mid-1, result)
	} else if arr[mid] < result {
		BinarySearch(arr, mid+1, right, result)
	} else {
		fmt.Println("result : \n", mid)
	}
}

func main() {
	//定义一个数组
	arr := []int{1, 2, 5, 7, 15, 25, 30, 36, 39, 51, 67, 78, 80, 82, 85, 91, 92, 97}
	//BinaryFind(&arr, 0, len(arr) - 1, 30)
	BinarySearch(arr, 0, len(arr)-1, 30)

}
