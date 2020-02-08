package main

import "fmt"

func main() {
	var numOfWorkers int
	fmt.Scanf("%d", &numOfWorkers)

	c := make(chan int)
	OutChanArr := make([]chan map[int]int, numOfWorkers)
	done := make(chan bool)
	collateChan := make(chan map[int]int)

	go genFunc(c)

	for s := 0; s < numOfWorkers; s++ {
		OutChanArr[s] = fact(c, done, s)
	}

	collateFunc(done, collateChan, OutChanArr...)

	go func() {
		for s := 0; s < numOfWorkers; s++ {
			<-done
		}
		close(collateChan)
	}()

	for i := range collateChan {
		fmt.Println(i)
	}
}

func genFunc(c chan<- int) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 20; j++ {
			c <- j
		}
	}
	close(c)
}

func fact(cin <-chan int, done chan bool, routineName int) chan map[int]int {
	cout := make(chan map[int]int)
	go func() {
		counter := 0
		for i := range cin {
			out := 1
			for j := i; j > 1; j-- {
				out *= j
			}
			cout <- map[int]int{i: out}
			counter++
		}
		fmt.Println("Goroutine fact ", routineName, "processed", counter, "items")
		close(cout)
		fmt.Println("Goroutine fact", routineName, "is finished")
	}()
	return cout
}

func collateFunc(done chan bool, collateChan chan map[int]int, c ...chan map[int]int) {
	for idx, ci := range c {
		go func(ci chan map[int]int, idx int) {
			counter := 0
			for i := range ci {
				collateChan <- i
				counter++
			}
			fmt.Println("Goroutine consume ", idx, "consumed", counter, "items")
			done <- true
		}(ci, idx)
	}
}
