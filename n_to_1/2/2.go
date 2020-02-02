package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	f := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		f <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		f <- true
	}()

	go func() {
		<-f
		<-f
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}
