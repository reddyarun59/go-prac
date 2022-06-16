package main

import "fmt"

func main() {
	fmt.Println("Variables")

	var i int = 42

	fmt.Println(i)

	fmt.Printf("Variable is of type %T \n", i)

	var username string = "Arun"
	fmt.Println(username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T \n", isLoggedIn)

	// default values and alies

	var anotherVar string
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type %T \n", anotherVar)

	// implicit type

	var something = "Hello World"
	fmt.Println(something)
	fmt.Printf("Variable is of type %T \n", something)
	// something = 3
	// fmt.Println(something)

	// no var style

	newVar := 32
	fmt.Println(newVar)
}
