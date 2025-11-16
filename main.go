package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	var address string
	var from int
	var to int
	var wg sync.WaitGroup
	fmt.Printf("ip地址:")
	fmt.Scanf("%v", &address)
	fmt.Printf("起始端口:")
	fmt.Scanf("%v", &from)
	fmt.Printf("结束端口:")
	fmt.Scanf("%v", &to)
	able_ports := make([]int, 0)

	results := make(chan int)
	fmt.Println("正在扫描")
	for i := from; i <= to; i++ {
		wg.Add(1)
		go check_port(address, i, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for p := range results {
		able_ports = append(able_ports, p)
	}

	fmt.Println("开放的端口:")
	for _, v := range able_ports {
		fmt.Println(v)
	}

}
func check_port(address string, port int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", address, port), 200*time.Millisecond)
	if err != nil {
		return
	}
	defer conn.Close()

	results <- port
}
