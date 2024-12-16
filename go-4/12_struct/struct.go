package main

import (
	"fmt"
	"time"
)

type order struct {
	id     string
	amount float32
	status string
	createdAt time.Time // time package
}

func newOrder(id string,amount float32,status string) *order {
	// initial setup goes here...
	myOrder := order  {
		id: id,
		amount: amount,
		status: status,
		createdAt: time.Now(),
	}
	return &myOrder
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}


func main() {
	language := struct {
		name string
		isGood bool
	} {"golang",true}

	fmt.Println(language)


	myOrder := order  {
		id: "1",
		amount: 50.00,
		status: "received",
		createdAt: time.Now(),
	}

	// myOrder := newOrder("1",30.50,"received")

	fmt.Println(myOrder)

	myOrder.changeStatus("confirmed")
	fmt.Println(myOrder.getAmount())

	myOrder.createdAt = time.Now()

	fmt.Println("Order Status :",myOrder.status)

	fmt.Println("order struct",myOrder)

	myOrder2 := order {
		id: "2",
		amount: 60.00,
		status: "pending",
		createdAt: time.Now(),
	}

	fmt.Println("order 2 struct",myOrder2)
}
