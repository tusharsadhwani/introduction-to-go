package main

import "fmt"

func avg(a, b int) float64 {
	ans := float64(a+b) / 2
	return ans
}

func main() {
	val1 := 5
	val2 := 7
	avgVal := avg(val1, val2)
	fmt.Println(avgVal)
	fmt.Println(avg(32, 52))

	greet := func() {
		fmt.Println("Hello Techies!")
	} // Inline / Anonymous Function
	greet()
}
