package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numChannel := make(chan int)
	// go processNum(numChannel)
	// numChannel <- 5
	go processing(numChannel)
	for {
		numChannel <- rand.Intn(100)
	}

	result := make(chan int)
	go sum(result, 4, 5)

}

func processNum(numChannel chan int) {
	fmt.Println("Processing... ", <-numChannel)
}

func processing(numChannel chan int) {
	for n := range numChannel {
		fmt.Println("Processing... ", n)
		time.Sleep(time.Second)
	}
}

func sum(result chan int, num1 int, num2 int) {
	r := num1 + num2
	result <- r
}
