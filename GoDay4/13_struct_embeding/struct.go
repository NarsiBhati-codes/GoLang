package main

import (
	"fmt"
	"time"
)


type customer struct {
	name string
	phone string
}

type order struct {
	id     string
	amount float32
	status string
	createdAt time.Time // nanosecond precision
	customer
}


func main() {

	newOrder := order {
		id: "1",
		amount: 40.00,
		status: "received",
		createdAt: time.Now(),
		customer: customer{
			name: "Narsi",
			phone: "454343434",
		},
	}

	newOrder.customer.name = "Mansi"

	fmt.Println(newOrder)
}