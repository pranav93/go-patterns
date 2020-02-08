package main

import (
	"fmt"
	"time"
)

func main() {
	var counter int
	counterChan := make(chan int)
	done := make(chan bool)

	go incrementor(counterChan, done, "foo")
	go incrementor(counterChan, done, "bar")

	go func() {
		for i := range counterChan {
			counter += i
		}
	}()

	<-done
	fmt.Println(counter)
	<-done
	fmt.Println(counter)
}

func incrementor(counterChan chan int, done chan bool, s string) {
	time.Sleep(time.Duration(time.Millisecond * 100))
	for i := 0; i < 1000; i++ {
		counterChan <- 1
	}
	done <- true
}
