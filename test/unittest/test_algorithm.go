package unittest

import "C"
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func Fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}

// SyncMux 并发请求API
func SyncMux(){
	var wg sync.WaitGroup
	var urls = []string{
		"https://api.github.com/",
	}
	for _, url := range urls {
		// 对 WaitGroup 计数器执行加一操作
		wg.Add(1)
		// 启动一个 goroutine ，用于获取给定的 URL
		go func(url string) {
			// 在 goroutine 执行完毕时，对计数器执行减一操作
			defer wg.Done()
			// 获取 URL
			resp,_ := http.Get(url)
			defer resp.Body.Close()
			b,_ := ioutil.ReadAll(resp.Body)
			fmt.Printf("%s\n",string(b))
		}(url)
	}

	// 等待直到所有 HTTP 获取操作都执行完毕为止
	wg.Wait()
}

