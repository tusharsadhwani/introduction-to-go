package main

import "fmt"

func incByValue(i int) { i++ }

func incByReference(i *int) { (*i)++ }

func main() {
	x := 41
	incByValue(x)
	fmt.Println("By value:", x)

	incByReference(&x)
	fmt.Println("By Ref:", x)
	fmt.Printf("Ptr: %v, Type: %T", &x, &x)
}
