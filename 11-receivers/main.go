package main

import "fmt"

type employeeCode string

type employee struct {
	name   string
	code   employeeCode
	salary int
}

func (e employee) verify() {
	if len(e.code) == 7 {
		fmt.Println("Employee Verified")
	} else {
		fmt.Println("Some Error occured")
	}
}

func main() {
	john := employee{"John Smith", "emp1001", 35000}
	john.verify()
}
