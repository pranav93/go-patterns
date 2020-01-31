package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int64
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go incrementor("foo")
	go incrementor("bar")
	wg.Wait()
}

func incrementor(s string) {
	for i := 0; i < 45; i++ {
		time.Sleep(time.Duration(time.Millisecond * 100))
		atomic.AddInt64(&x, 1)
		fmt.Println(s, x)
	}
	wg.Done()
}
