package main

import "fmt"

func main() {
	// String
	var username string = "rashid"
	fmt.Println("Variables", username)
	fmt.Printf("Variables type is: %T \n", username)

	// Boolean
	var isLoggedIn bool = false
	fmt.Println("Variables", isLoggedIn)
	fmt.Printf("Variables type is: %T \n", isLoggedIn)

	// int
	var number int = 255
	fmt.Println("Variables", number)
	fmt.Printf("Variables type is: %T \n", number)

	// unit
	var smallVal uint8 = 255
	fmt.Println("Variables", smallVal)
	fmt.Printf("Variables type is: %T \n", smallVal)

	// float
	var smallFloat float32 = 255.00
	fmt.Println("Variables", smallFloat)
	fmt.Printf("Variables type is: %T \n", smallFloat)

	// default values and some aliases for int
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variables type is: %T \n", anotherVariable)

	// default values and some aliases for bool
	var anotherVariableBool bool
	fmt.Println(anotherVariableBool)
	fmt.Printf("Variables type is: %T \n", anotherVariableBool)

	// no var style
	numberOfSubscribers := 2000
	fmt.Println(numberOfSubscribers)
	fmt.Printf("Variables type is: %T \n", numberOfSubscribers)

}
