package main

import (
	"errors"
	"fmt"
)

func test(a string) (int, error) {
	if len(a) == 0 {
		return 0,
			errors.New("string is empty")
	}
	return len(a), nil
}

func main() {
	result, err := test("")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
