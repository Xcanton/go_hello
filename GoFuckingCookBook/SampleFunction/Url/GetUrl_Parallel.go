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
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.0fs elapsed \n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {

	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // 发送到通道ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 不要泄露资源

	if err != nil {
		ch <- fmt.Sprintf("while reading %s:\t%v", url, err)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", secs, nbytes, url)
}
