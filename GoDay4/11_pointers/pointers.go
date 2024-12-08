package main

import "fmt"

// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In change num",num)
// }

// by reference
func changNum(num *int) {
	*num = 5
	fmt.Println("In change num",*num)
}

func main() {
	num := 1 

	// changeNum(num)
	// fmt.Println("Memory address :",&num)

	changNum(&num)

	fmt.Println("After changeNum in main ",num)

}