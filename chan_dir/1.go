package main

import "fmt"

func main() {
	sc := make(chan int)
	c := incrementor()
	puller(c, sc)
	fmt.Println("sum is", <-sc)
}

func incrementor() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func puller(c <-chan int, sc chan<- int) {
	go func() {
		sum := 0
		for i := range c {
			sum += i
		}
		sc <- sum
		close(sc)
	}()
}
