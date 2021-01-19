package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ntime := time.Now().UnixNano()
	fmt.Println(ntime)
	fmt.Println(len(strconv.FormatInt(ntime, 10)))
	fmt.Println(len("1610553600000000000"))
	fmt.Println(len("1610171300911827873"))

	now := time.Now().UnixNano()
	fmt.Println(now)

	beforeWeekDay := time.Now().AddDate(0, 0, -30).UnixNano()
	fmt.Println(beforeWeekDay)

}
