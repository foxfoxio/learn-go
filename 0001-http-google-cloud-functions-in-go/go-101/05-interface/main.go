package main

import "fmt"

type animal interface{ makeSound() }

type dog struct{}
type cat struct{}

func (d dog) makeSound() {
	fmt.Println("i say woof!!")
}

func (c cat) makeSound() {
	fmt.Println("i say meow!!")
}

func makeSound(a animal) {
	a.makeSound()
}

func main() {
	makeSound(dog{})
	makeSound(cat{})
}
