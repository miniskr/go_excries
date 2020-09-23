package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("sunnim.com:%d", p)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 1000)
	results := make(chan int)

	var openPorts []int
	var closePorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i < 9000; i++ {
			ports <- i
		}
	}()

	for i := 1; i < 9000; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		} else {
			closePorts = append(closePorts, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)
	sort.Ints(closePorts)

	for _, port := range openPorts {
		fmt.Printf("%d 打开了 \n", port)
	}
}
