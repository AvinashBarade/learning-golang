package main

import "fmt"

func compute(x, y int, fn func(int, int) int) int {
	return fn(x, y)
}

func main() {

	add := func(x, y int) int {
		return (x + y)
	}

	sub := func(x, y int) int {
		return x - y
	}

	fmt.Println("Sum=", compute(5, 6, add))
	fmt.Println("Div=", compute(5, 6, sub))
}
