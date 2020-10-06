package main

import "fmt"

func main() {
	// for loop
	for count := 0; count <= 10; count++ {
		fmt.Println("My counter is at", count)
	}

	// for-range
	entry := []string{"john", "doe"}
	for i, val := range entry {
		fmt.Printf("index %d, value %s\n", i, val)
	}

	day := "sunday"

	if day == "sunday" || day == "saturday" {
		fmt.Println("rest")
	} else if day == "monday" {
		fmt.Println("groan")
	} else {
		fmt.Println("work")
	}

	// switch case
	switch day {
	case "sunday":
		fallthrough
	case "saturday":
		fmt.Println("rest")
		break
	case "monday":
		fmt.Println("groan")
		break
	default:
		fmt.Println("work")
	}
}
