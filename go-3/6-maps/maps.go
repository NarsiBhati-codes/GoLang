package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating map 

	m := make(map[string]string) 

	// setting an element 
	m["name"] = "golang"
	m["area"] = "backend"

	// get element 
	fmt.Println(m["name"])
	fmt.Println(m["area"])

	// IMP: if key does not exists in the map then it returns zero value
	fmt.Println(m["something"])

	// with int values 
	m2 := make(map[string]int)
	m2["age"]  = 24 

	fmt.Println(m2["age"])
	fmt.Println(m2["phone"])

	// length
	fmt.Println(len(m))

	// delete 
	delete(m, "area")
	fmt.Println(m["area"])

	// empty the map
	clear(m)
	fmt.Println(m)

	// without using make function 
	m3 := map[string]int{"price": 40 ,"phone": 3}
	fmt.Println(m3)


	// find element in map 
	value, ok := m3["price"]

	fmt.Println(value)

	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}


	// maps package 
	m4 := map[string]int{"price": 40 ,"phone": 3}
	m5 := map[string]int{"price": 40 ,"phone": 3}

	fmt.Println(maps.Equal(m4,m5))
}