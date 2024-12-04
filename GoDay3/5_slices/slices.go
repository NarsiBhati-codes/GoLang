package main

import (
	"fmt"
	"slices"
)

// slices -> dynamic
// most used construct in go
// + useful methods

func main() {
	// uninitialized slices is nil
	var nums []int 
	fmt.Println(nums == nil)
	
	fmt.Println(len(nums))

	var nums2 = make([]int, 0, 5) //  make( type, initial size, capacity)

	fmt.Println(nums2 == nil)
	fmt.Println(nums2)

	// append -> add value into slice
	nums2 = append(nums2,1)  
	nums2 = append(nums2,2)
	
	fmt.Println(nums2)
	fmt.Println(len(nums2))
	fmt.Println(cap(nums2))  // cap -> capacity


	// give values using indexes 
	var nums3 = make([]int, 2, 5)
	fmt.Println("Before :",nums3)
	nums3[0] = 1
	nums3[1] = 2
	fmt.Println("After :",nums3)


	//copy function 
	var nums4 = make([]int,0,5)
	nums4 = append(nums4, 1)

	var nums5 = make([]int,len(nums4))
	fmt.Println("Before copy :",nums4,nums5)

	copy(nums5, nums4) // copy(dst, src)
	fmt.Println("After copy :",nums4,nums5)

	// slice operator 
	var nums6 = []int{1,2,3}
	fmt.Println(nums6[0:2])
	fmt.Println(nums6[:2])
	fmt.Println(nums6[1:])

	// slices package
	var nums7 = []int{1,2} 
	var nums8 = []int{1,2} 

	fmt.Println(slices.Equal(nums7,nums8)) // return bool, compare both values 

}