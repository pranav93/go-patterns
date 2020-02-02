package main

import (
	"fmt"
)

func main() {
	n := 10
	c := make(chan int)
	f := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			f <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-f
		}
		close(c)
	}()

	count := 0
	for i := range c {
		fmt.Println(i)
		count++
	}
	fmt.Println("count", count)
}
