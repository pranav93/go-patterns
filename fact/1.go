package main

import (
	"fmt"
)

func main() {
	var factorialNum int
	fmt.Printf("Enter a number : ")
	fmt.Scanf("%d\n", &factorialNum)

	if factorialNum <= 0 {
		fmt.Println("Wrong input")
		return
	}

	chGen := make(chan int)

	factorialGenerator(factorialNum, chGen)
	chRes := factorialConsumer(chGen)
	fmt.Println("factorial", <-chRes)
}

func factorialGenerator(factorialNum int, chGen chan<- int) {
	go func() {
		for i := 2; i <= factorialNum; i++ {
			chGen <- i
		}
		close(chGen)
	}()
}

func factorialConsumer(chGen <-chan int) <-chan int {
	chRes := make(chan int)
	go func() {
		sum := 1
		for i := range chGen {
			sum *= i
		}
		chRes <- sum
		close(chRes)
	}()
	return chRes
}
