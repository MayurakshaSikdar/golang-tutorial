package main

import (
	"fmt"
	"strconv"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

func main() {

	// Pointers
	p := 1
	fmt.Println("Memory => ", &p)
	fmt.Println("Value => ", p)
	changeNum(&p)
	fmt.Println("New Value => ", p)

	// Struct
	order := order{
		id:     strconv.Itoa(time.Now().Nanosecond()),
		amount: 100.0,
		status: "completed",
	}
	order.createdAt = time.Now()
	fmt.Println(order)
	fmt.Println(order.createdAt.String())

	order.changeStatus("archive")
	fmt.Println(order)

	myOrder := NewOrder(strconv.Itoa(time.Now().Nanosecond()), 200.00, "pending")
	fmt.Println(*myOrder)

	language := struct {
		name   string
		isGood bool
	}{name: "English", isGood: true}
	fmt.Println(language)
}

func changeNum(num *int) {
	*num = 5
	fmt.Println("Change Num => ", *num)
}

func (o *order) changeStatus(status string) {
	o.status = status
}

func NewOrder(id string, amount float32, status string) *order {
	order := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &order
}
