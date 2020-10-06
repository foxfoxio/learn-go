package main

import "fmt"

var messages = make(chan int)

func producer(n int) {
	for i := 0; i < n; i++ {
		messages <- i
	}
}

func consumer() {
	for {
		fmt.Println(<-messages)
	}
}

func main() {
	go consumer()
	go producer(10)

	fmt.Println("press enter when done")
	var input string
	fmt.Scanln(&input)
}
