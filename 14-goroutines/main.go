package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")
	go f("goroutine")

	go func(msg string) {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(msg)
	}("another goroutine")

	time.Sleep(time.Second * 3)
	fmt.Println("done")
}
