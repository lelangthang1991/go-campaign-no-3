package main

import (
	"fmt"
)

func test(i int) string {
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
		return "one"
	case 2:
		fmt.Println("two")
		return "two"
	case 3:
		fmt.Println("three")
		return "three"
	}
	return "nothing"
}
