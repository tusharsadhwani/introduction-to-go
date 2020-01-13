package main

import "fmt"

func main() {
	var num int
	num = 42
	var num2 = 42 // type inferred
	num3 := 42    // shorter form

	fmt.Println(num)
	fmt.Println(num2)
	fmt.Println(num3)
}
