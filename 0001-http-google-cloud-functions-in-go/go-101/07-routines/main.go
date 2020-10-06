package main

import "fmt"
import "../02-package/math"

func f(n int) {
	fmt.Printf("%v! = %v\n",
		n, math.Fact(n))
}

func main() {
	fmt.Println(
		"press enter when done")

	go f(20)
	go f(5)
	go f(0)

	var input string
	fmt.Scanln(&input)
}
