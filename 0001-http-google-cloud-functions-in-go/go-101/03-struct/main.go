package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var bob = person{
		name: "Bob",
		age:  26,
	}
	fmt.Println(bob)

	bob.age = 27
	fmt.Println(bob.age)
}
