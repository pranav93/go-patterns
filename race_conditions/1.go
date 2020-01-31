package main

import (
	"fmt"
	"sync"
	"time"
)

var x int
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("foo")
	go incrementor("bar")
	wg.Wait()
}

func incrementor(s string) {
	for i := 0; i < 45; i++ {
		time.Sleep(time.Duration(time.Millisecond * 100))
		mutex.Lock()
		counter := x
		counter++
		fmt.Println(s, x)
		x = counter
		mutex.Unlock()
	}
	wg.Done()
}
