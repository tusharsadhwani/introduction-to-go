package main

import "fmt"

func main() {
	arr1 := [3]int{5, 6, 7}   // Array of length 3
	arr2 := [...]int{5, 6, 7} // Alternative syntax

	arr1[0] = 500
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)

	slice1 := []int{1, 2, 3, 4}    // Slices don't have fixed length
	slice2 := append(slice1, 5, 6) // Doesn't affect the original slice
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
}
