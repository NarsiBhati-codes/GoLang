package main

import "fmt"

func main() {
	fmt.Println(`main file is here call other files`)

	fmt.Println(`file 1`)
	helloWorld()

	fmt.Println(`file 2`)
	simpleValues()

	fmt.Println(`file 3`)
	variables()
}