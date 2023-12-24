package main

import (
	"fmt"
	"strings"
)

func toCN(num int) string {
	//1、数字为0
	if num == 0 {
		return "零"
	}
	var ans string
	//数字
	szdw := []string{"十", "百", "千", "万", "十万", "百万", "千万", "亿"}
	//数字单位
	sz := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	res := make([]string, 0)

	//数字单位角标
	idx := -1
	for num > 0 {
		//当前位数的值
		x := num % 10
		//2、数字大于等于10
		// 插入数字单位，只有当数字单位角标在范围内，且当前数字不为0 时才有效
		if idx >= 0 && idx < len(szdw) && x != 0 {
			res = append([]string{szdw[idx]}, res...)
		}
		//3、数字中间有多个0
		// 当前数字为0，且后一位也为0 时，为避免重复删除一个零文字
		if x == 0 && len(res) != 0 && res[0] == "零" {
			res = res[1:]
		}
		// 插入数字文字
		res = append([]string{sz[x]}, res...)
		num /= 10
		idx++
	}
	//4、个位数为0
	if len(res) > 1 && res[len(res)-1] == "零" {
		res = res[:len(res)-1]
	}
	//合并字符串
	for i := 0; i < len(res); i++ {
		ans = ans + res[i]
	}
	return ans
}

func transfer(num int) string {
	ori := num
	chineseMap := []string{"圆整", "十", "百", "千", "万", "十", "百", "千", "亿", "十", "百", "千"}
	//chineseNum := []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
	chineseNum := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}

	listNum := []int{}
	for ; num > 0; num = num / 10 {
		listNum = append(listNum, num%10)
	}
	n := len(listNum)
	chinese := ""
	//注意这里是倒序的
	for i := n - 1; i >= 0; i-- {
		fmt.Println("chinese", chinese)
		chinese = fmt.Sprintf("%s%s%s", chinese, chineseNum[listNum[i]], chineseMap[i])
	}
	fmt.Println("listNum", ori, listNum, chinese)
	//注意替换顺序
	for {
		copychinese := chinese
		copychinese = strings.Replace(copychinese, "零万", "万", 1)
		copychinese = strings.Replace(copychinese, "零亿", "亿", 1)
		copychinese = strings.Replace(copychinese, "零十", "零", 1)
		copychinese = strings.Replace(copychinese, "零百", "零", 1)
		copychinese = strings.Replace(copychinese, "零千", "零", 1)
		copychinese = strings.Replace(copychinese, "零零", "零", 1)
		copychinese = strings.Replace(copychinese, "零圆", "圆", 1)

		if copychinese == chinese {
			break
		} else {
			chinese = copychinese
		}
	}

	return chinese
}

func main() {
	num := 1020201
	chineseRepresentation := toCN(num)
	fmt.Printf("%d 转换成中文是：%s\n", num, chineseRepresentation)

	fmt.Println(1024, transfer(1024))
	fmt.Println(1004, transfer(1004))
	fmt.Println(101004, transfer(101024))
	fmt.Println(1010004, transfer(1010004))
	fmt.Println(3000010004, transfer(3000010004))
}
