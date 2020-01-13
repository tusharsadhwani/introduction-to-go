package main

import "fmt"

type book struct {
	name string
	code int
}

func main() {
	book1 := book{"Alice in Wonderland", 1301}
	fmt.Println(book1)
	fmt.Printf("%v\n", book1)
	fmt.Printf("%#v\n", book1)
}
