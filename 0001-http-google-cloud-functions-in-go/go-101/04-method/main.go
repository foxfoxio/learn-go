package main

import "fmt"

type rect struct {
	width, height int
}

// pass by value
func (r rect) area() int {
	return r.width * r.height
}

// pass by reference
func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	var r = rect{width: 20, height: 10}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
}
