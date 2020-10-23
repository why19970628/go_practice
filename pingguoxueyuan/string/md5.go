package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// MD5 方法

func getTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func getPid()  {
	pid := os.Getpid()
	fmt.Printf("进程 PID: %d \n", pid)

	prc := exec.Command("ps", "-p", strconv.Itoa(pid), "-v")
	out, err := prc.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
func main() {
	str := "12345"
	fmt.Printf("MD5(%s): %s\n", str, MD5(str))

	fmt.Printf("current time str : %s\n", getTimeStr())

	// 线程id
	fmt.Println(" Getppid", unix.Getppid())
	getPid()
}

