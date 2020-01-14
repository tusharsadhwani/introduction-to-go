package main

import "fmt"

func test(x func()) {
	fmt.Println("running the function")
	x()
	fmt.Printf("%T", x)
}

func main() {
	x := func() { fmt.Println("test") }
	test(x)
}
