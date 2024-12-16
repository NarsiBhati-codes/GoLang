package main

import (
	"fmt"
	"time"
)

func main() {

	// simple switch
	i := 5

	switch i {
		case 1:
			fmt.Println("one")
		case 2: 
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		default: 
			fmt.Println("other")
	}


	// multiple condition switch

	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday: 
			fmt.Println("it's weekend")
		default: 
			fmt.Println("it's workday")
	}

	// type switch 

	whoAmI := func( i interface{}){
		switch t := i.(type) {
			case int:
				fmt.Println("it's a integer")
			case string:
				fmt.Println("it's a string")
			case bool:
				fmt.Println("it's a boolean")
			default:
				fmt.Println("other",t)
			
		}
	}

	whoAmI(55)
}