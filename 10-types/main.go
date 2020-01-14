package main

import "fmt"

type employeeCode string

type employee struct {
	name   string
	code   employeeCode
	salary int
}

func main() {
	john := employee{
		"John Smith",
		"emp1001",
		35000,
	}
	fmt.Printf("%#v", john)
}
