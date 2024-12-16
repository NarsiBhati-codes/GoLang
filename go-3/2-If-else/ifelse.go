package main

import "fmt"

func main() {

	//  if else 
	age := 18

	if age >= 18 {
		fmt.Println("person is adult")
	} else {
		fmt.Println("person is not adult")
	}


	// if else-if else
	otherAge := 16 

	if otherAge >= 18 {
		fmt.Println("person is adult")
	} else if otherAge >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is kid")
	}
	

	// logical operator

	var role = "admin"
	var hasPermissions = false 

	if role == "admin" && hasPermissions {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}


	// we can declare a variable inside if construct

	if age := 15; age >=18 {
		fmt.Println("person is an adult",age)
	} else if age >=12 {
		fmt.Println("person is teenager",age)
	}

	// Go does not have ternary operator, you will have to use normal if else 

}
