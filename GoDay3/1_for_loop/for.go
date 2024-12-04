package main

import "fmt"

func main() {

	// for loop like while loop
	i := 1
	for i <= 100 {
		fmt.Println(i)

		i = i + 1
	}

	// infinite for loop
	// for {
		// fmt.Println()
	// }

	for i:= 1; i<=10; i++ {
		// there if break and continue keyword you can use in loop
	}

	// range in loop 
	for i:= range(10) {
		fmt.Println(i)
	}


}