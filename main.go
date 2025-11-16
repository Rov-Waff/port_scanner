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
	var mu sync.Mutex
	fmt.Printf("ip地址:")
	fmt.Scanf("%v", &address)
	fmt.Printf("起始端口:")
	fmt.Scanf("%v", &from)
	fmt.Printf("结束端口:")
	fmt.Scanf("%v", &to)
	able_ports := make([]int, 0)
	fmt.Printf("开始扫描\n")
	for i := from; i <= to; i++ {
		wg.Add(1)
		go check_port(address, i, &able_ports, &wg, &mu)
	}
	wg.Wait()
	fmt.Println("开放的端口:")
	for _, v := range able_ports {
		fmt.Println(v)
	}
}

func check_port(address string, port int, sl *[]int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", address, port), 200*time.Millisecond)
	if err != nil {
		return
	}
	defer conn.Close()

	mu.Lock()
	*sl = append(*sl, port)
	mu.Unlock()
}
