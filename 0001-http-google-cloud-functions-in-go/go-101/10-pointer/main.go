package main

import "fmt"

func getPointer() *int {
	a := 234
	return &a
}

func main() {
	b := *getPointer()
	fmt.Println("Value is", b)

	a := new(int)
	*a = 456
	fmt.Println("Value is", a)
	fmt.Println("Value is", *a)
}
