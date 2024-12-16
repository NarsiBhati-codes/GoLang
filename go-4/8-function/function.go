package main

import "fmt"

func add(a int, b int) int {
	return ( a + b )
}

func getLanguages() (string, string, bool) {
	return "go lang","JavaScript", true
}

func processIt(fn func(a int) int) {
	fmt.Println(fn(1))
}

func returnFunc() func(a int) int {
	return func(a int) int {
		return 5
	}
}

func main() {
    
	// normal function
	result := add(10,5)
	fmt.Println(result)

	// multi value return function
	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1,lang2)


	//  function as a parameter
	fn := func(a int) int {
		return 2
	}

	processIt(fn)


	// return function 
	fn2 := returnFunc()
	fmt.Println(fn2(3))

}