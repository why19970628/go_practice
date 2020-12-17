package main


// goroutine 通过信号协作退出
import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

func cancelByChannel2(c net.Conn, quit <-chan time.Time, wg *sync.WaitGroup) {
	defer c.Close()
	defer wg.Done()

	for {

		select {
		case <-quit:
			fmt.Println("cancel goroutine by channel!")
			return
		default:
			_, err := io.WriteString(c, "hello cancelByChannel")
			if err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	quit := time.After(time.Second * 10)
	go cancelByChannel2(conn, quit, &wg)
	wg.Wait()
	// mac: nc localhost 8000
}