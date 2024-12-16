package main

import "fmt"


func sum(nums ...int) int {
	var sum int 
	for i:= range nums {
		sum = sum + i
	}
	return sum 
}

func main() {
	fmt.Println(sum(1,2,3,3,4,4,5))


	nums := []int{1,2,3,4,5} 
	fmt.Println(sum(nums...))
}