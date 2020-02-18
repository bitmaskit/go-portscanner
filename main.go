package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
	"time"
)

var (
	addr    string
	from    int
	to      int
	wCnt    int
	results = make(chan int)
)

func worker(addr string, ports chan int) {
	var address string
	for p := range ports {
		address = fmt.Sprintf("%s:%d", addr, p)
		conn, err := net.DialTimeout("tcp", address, 5*time.Second)
		if err != nil {
			results <- 0
			continue
		} else {
			_ = conn.Close()
			results <- p
		}
	}
}

func main() {
	flag.StringVar(&addr, "addr", "localhost", "'addr' indicates address you want to scan.")
	flag.IntVar(&from, "from", 1, "'from' indicates starting port.")
	flag.IntVar(&to, "to", 65535, "'to' indicates ending port.")
	flag.IntVar(&wCnt, "w", 100, "'w' indicates workers count. The more workers you have the faster scan will be. But ")
	flag.Parse()

	fmt.Printf("Creating %d workers to scan ports.\n", wCnt)

	ports := make(chan int, wCnt)
	for i := 0; i < wCnt; i++ {
		go worker(addr, ports)
	}
	fmt.Printf("%d workers to scan ports.\n", wCnt)
	go func() {
		fmt.Printf("Scanning %s from %d to %d\n", addr, from, to)
		for i := from; i <= to; i++ {
			ports <- i
		}
	}()

	var openPorts []int
	for i := from; i < to; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d opened \n", port)
	}
	close(ports)
	close(results)
}
