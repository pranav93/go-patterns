package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		wg.Done()
	}()

	wg.Wait()
}
