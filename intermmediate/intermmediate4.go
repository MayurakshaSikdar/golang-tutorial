package main

import (
	"fmt"
	"sync"
	"time"
)

type OrderStatus int
type OrderStatus2 string

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

const (
	Received2  OrderStatus2 = "Received"
	Confirmed2              = "Confirmed"
	Prepared2               = "Prepared"
	Delivered2              = "Delivered"
)

func main() {
	// Enums
	changeOrderStatus(Confirmed)
	changeOrderStatus2(Delivered2)

	// Generics
	printSlice([]int{1, 2, 3, 4})
	printSlice([]string{"One", "Two", "Three", "Four"})
	printSliceScoped([]string{"One", "Two", "Three", "Four"})
	printSliceComparable([]int{1, 2, 3, 4})

	// Go Routines
	for i := 1; i <= 10; i++ {
		go task(i)
	}
	time.Sleep(time.Second * 2)

	// Go Routines -- Wait Groups
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go task2(i, &wg)
	}
	wg.Wait()
}

func changeOrderStatus(status OrderStatus) {
	fmt.Println("### Inside `changeOrderStatus`... ###")
	fmt.Println("Order Status", status)
}
func changeOrderStatus2(status OrderStatus2) {
	fmt.Println("### Inside `changeOrderStatus2`... ###")
	fmt.Println("Order Status", status)
}

func printSlice[T any](items []T) {
	fmt.Println("### Inside `printSlice`... ###")
	for _, item := range items {
		fmt.Println(item)
	}
}

func printSliceScoped[T int | string](items []T) {
	fmt.Println("### Inside `printSliceScoped`... ###")
	for _, item := range items {
		fmt.Println(item)
	}
}

func printSliceComparable[T comparable](items []T) {
	fmt.Println("### Inside `printSliceComparable`... ###")
	for _, item := range items {
		fmt.Println(item)
	}
}

func task(id int) {
	fmt.Println("Task => ", id)
}

func task2(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Task => ", id)
}
