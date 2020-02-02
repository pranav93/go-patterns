package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	f := make(chan bool)

	go func() {
		for i := 0; i < 1000; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		count := 0
		for i := range c {
			fmt.Println(i)
			count++
		}
		fmt.Println("count1 ->", count)
		f <- true
	}()
	go func() {
		count := 0
		for i := range c {
			fmt.Println(i)
			count++
		}
		fmt.Println("count2 ->", count)
		f <- true
	}()

	<-f
	<-f
}
