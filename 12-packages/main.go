package main

import (
	"fmt"

	"github.com/tusharsadhwani/introduction-to-go/12-packages/util"
)

func main() {
	randomDigit := util.RandomDigit()
	fmt.Println(randomDigit)
}
