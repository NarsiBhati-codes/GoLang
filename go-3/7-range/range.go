package main

import "fmt"

func main() {

	// range with slice 
	nums := []int{6,7,8}

	for index, value := range nums {
		fmt.Println(value,index)	
	}

	// range with maps
	m := map[string]int{"price": 40,"phone": 30}

	for key, value := range m {
		fmt.Println(key,value)
	}

	// range with string 
	// return unicode point rune
	// starting byte of rune
	// 255 -> 1 byte 
	for i, c := range "go lang" {
		fmt.Println("on index [",i,"] :",string(c))
	}
}