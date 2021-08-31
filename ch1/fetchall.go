// 并发获取 URL 并报告它们的时间和大小
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // 启动一个 goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // 从通道 ch 接收
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // 发送到通道 ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.body)
	resp.body.close()
	if err != nul {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
