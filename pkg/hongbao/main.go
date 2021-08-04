package main

import (
	"fmt"
	"math/rand"
	"time"
)

// DoubleRandom 二次随机算法
func DoubleRandom(count, amount int64) int64 {
	if count == 1 {
		return amount
	}
	min := int64(1)

	max := amount - min*count
	rand.Seed(time.Now().UnixNano())
	// 1次随机，计算出一个种子金额作为基数
	seed := rand.Int63n(count*2) + 1
	n := max/seed + 1
	// 2次随机，计算红包金额序列元素
	x := rand.Int63n(n)
	return x + min
}

//二倍均值算法,count剩余个数,amount剩余金额
// DoubleAverage
func DoubleAverage(count, amount int64) int64 {
	//最小钱
	min := int64(1)

	if count == 1 {
		//返回剩余金额
		return amount
	}

	//计算最大可用金额,min最小是1分钱,减去的min,下面会加上,避免出现0分钱
	max := amount - min*count
	//计算最大可用平均值
	avg := max / count
	//二倍均值基础加上最小金额,防止0出现,作为上限
	avg2 := 2*avg + min
	//随机红包金额序列元素,把二倍均值作为随机的最大数
	rand.Seed(time.Now().UnixNano())
	//加min是为了避免出现0值,上面也减去了min
	x := rand.Int63n(avg2) + min

	return x
}

func main() {
	//初始10个红包, 10000分 = 100元钱
	count, amount := int64(9), int64(10000)
	//剩余金额
	remain := amount
	//验证红包算法的总金额,最后sum应该==amount
	sum := int64(0)
	//进行发红包
	for i := int64(0); i < count; i++ {
		x := DoubleAverage(count-i, remain)
		//fmt.Println(x)
		//金额减去
		remain -= x
		//发了多少钱
		sum += x
		//金额转成元
		fmt.Println(i+1, "=", float64(x)/float64(100))
	}
	fmt.Println()
	fmt.Println("总和 ", sum)
}
