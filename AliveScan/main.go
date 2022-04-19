package main

import (
	"AliveScan/coo"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	host   string
	thread int
	wg     sync.WaitGroup
)

func init() {
	flag.StringVar(&host, "h", "", "eg:-h 192.168.1.1/24")
	flag.IntVar(&thread, "t", 200, "eg:-t 100")
	flag.Parse()

	if host == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

}

func main() {
	start := time.Now()
	target := coo.ParseIP(host) //返回需要探测的ip列表
	ipCh := make(chan string, len(target))
	for _, s := range target {
		ipCh <- s
	}
	close(ipCh)
	for i := 0; i < thread; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ch := range ipCh {
				coo.Icmp(ch)
			}
		}()
	}
	wg.Wait()
	fmt.Println("用时", time.Since(start))

}
