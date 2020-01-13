package main

import "fmt"

func main() {
	var attendees int         // int8, int16, int32, int64
	var statusCode uint = 404 // uint8, uint16, uint32, uint64
	var b1, b2 bool = true, false
	var myName string = "Tushar"
	var char1 rune = 'S' // character type

	fmt.Println(attendees)
	fmt.Println(statusCode)
	fmt.Println(b1, b2)
	fmt.Println(myName, char1)
}
