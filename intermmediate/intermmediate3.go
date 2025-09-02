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

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway for testing purpose")
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string) {}

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

	// Interface
	newPaymentGw := paypal{}
	newPay := payment{
		gateway: newPaymentGw,
	}
	newPay.makePayment(100)
}

func changeNum(num *int) {
	fmt.Println("### Inside `changeNum`... ###")
	*num = 5
	fmt.Println("Change Num => ", *num)
}

func (o *order) changeStatus(status string) {
	fmt.Println("### Inside `changeStatus`... ###")
	o.status = status
}

func NewOrder(id string, amount float32, status string) *order {
	fmt.Println("### Inside `NewOrder`... ###")
	order := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &order
}

func (p payment) pay(amount float32) {
	fmt.Println("### Inside `pay`... ###")
	p.gateway.pay(amount)
}
