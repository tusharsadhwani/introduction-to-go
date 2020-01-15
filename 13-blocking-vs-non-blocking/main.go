package main

import (
	"fmt"
	"math/rand"
	"time"
)

func computeResult() {
	time.Sleep(2 * time.Second)
	result := rand.Intn(100)
	fmt.Println("Result:", result)
}

func main() {
	fmt.Println("Running blocking call...")
	computeResult()
	fmt.Println("This prints after 2 seconds")

	fmt.Println("Running non-blocking call...")
	go computeResult()
	fmt.Println("This prints immediately after")
	time.Sleep(time.Second * 5)
	fmt.Println("Exit program")
}
