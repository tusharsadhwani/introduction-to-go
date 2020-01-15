package main

import (
	"fmt"
	"time"
)

func squared(in chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * i
		}
	}()
	return out
}

func main() {
	nums := make(chan int, 5)
	out := squared(nums)
	nums <- 5
	nums <- 7
	nums <- 13

	fmt.Println(<-out)
	time.Sleep(time.Second)
	fmt.Println(<-out)
	time.Sleep(time.Second)
	fmt.Println(<-out)
}
