package main

import (
	"errors"
	"fmt"
)

func mustNotStringEmpty(s string) {
	if len(s) == 0 {
		panic(errors.New("string is empty"))
	}
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panic handled: %s\n", e)
		}
	}()

	mustNotStringEmpty("")
}
