package main

import "fmt"

func main() {
	// use array 
	// - fixed size, that is predictable
	// - memory optimization
	// - constant time access

	var nums [4]int

	fmt.Println("Array length :",len(nums))   // length of array
	
	nums[0] = 1
	fmt.Println("Zeroed value :",nums[0]) 
	fmt.Println("All values of Array :",nums)


	// do declare it in single line
	nums2 := [3]int{1,2,3} 
	fmt.Println(nums2)

	// 2D arrays 
	nums3 := [2][2]int{{1,2},{1,2}}
	for i:= 0;i<2;i++ {
		for j:=0;j<2;j++ {
			fmt.Println(nums3[i][j])
		}
	}

}