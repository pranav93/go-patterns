package main

import "fmt"

func main() {

	c := make(chan map[int]int)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			go fact(j, c)
		}
	}

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
}

func fact(i int, c chan<- map[int]int) {
	out := 1
	for j := i; j > 1; j-- {
		out *= j
	}
	c <- map[int]int{i: out}
}
