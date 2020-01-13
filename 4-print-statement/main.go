package main

import "fmt"

func main() {
	var isGolangBoring bool // all variables have a default "zero value"
	magicNumber := 42
	fullName := "Tushar Sadhwani"

	fmt.Println(isGolangBoring)

	fmt.Printf("%v\n", magicNumber)
	fmt.Printf("%T\n", magicNumber)

	fmt.Printf("My name is %v, ", fullName)
	fmt.Printf("this is a %T\n", fullName)

	fmt.Print("->") // Print doesn't add a newline by default
}
