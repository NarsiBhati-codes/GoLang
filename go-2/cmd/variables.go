package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

// datatype with ( var and const )

   var intNum int = 31   // int - int, int16, int32, int64
   fmt.Println(intNum)

   var floatNum float32 = 1212222.6 // float - float32, float64
   fmt.Println(floatNum)

// to use arithmetic operator but can't use with mix datatype like ( int + float ) isNotAllowed but  (int + int ) isAllowed 

// but if you want to do ( float + int ) you can do this 
	var intNum32 int32 = 2 
	var floatNum32 float32 = 10.1 
	var result float32 = floatNum32 + float32(intNum32) // conversion of int32 to float32 
	fmt.Println(result)


	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)  // divide 2 numbers
	fmt.Println(intNum1%intNum2) // get reminder



	// String 

	var myString1 string = "hello, world"
	fmt.Println("String 1 : " + myString1)

	var myString2 string = "hello" + " " + "world"
	fmt.Println("String 2 : " + myString2)

	var myString3 string = `hello
	world`
	fmt.Println("String 4 : " + myString3)

	var myString4 string = "hello \nworld"
	fmt.Println("String 3 : " + myString4)

	// length of string use 
	fmt.Println(utf8.RuneCountInString(myString4))

	
	// rune 
	var myRune rune = 'a'
	fmt.Println(myRune)

	// Boolean 
	var myBoolean bool = false 
	fmt.Println(myBoolean)

	// default values if we don't assign any all numbers like ( int - int, int16, int32, int64, float , rune ) it's 0 for string if will  be '' 

	var intNum3 rune 
	fmt.Println(intNum3)
	
	// shorthand
	myVar := "text"
	fmt.Println(myVar)

	var1 , var2 := 1, 2 // multiple assigns 
	fmt.Println(var1, var2)


	// const
	const myConst string = "const value"
	fmt.Println(myConst)
	
    
}