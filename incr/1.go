package main

import "fmt"

func main() {
	c1 := incrementor("foo")
	c2 := incrementor("bar")
	sc1 := puller(c1)
	sc2 := puller(c2)
	fmt.Println("sum is", <-sc1, <-sc2)
}

func incrementor(name string) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println(name, i)
		}
		close(c)
	}()
	return c
}

func puller(c <-chan int) <-chan int {
	sc := make(chan int)
	go func() {
		sum := 0
		for i := range c {
			sum += i
		}
		sc <- sum
		close(sc)
	}()
	return sc
}
