package main

import "fmt"

func main() {
	// function that resturns closure
	increment := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}
	//create a closure
	next := increment()

	fmt.Println(next())
	fmt.Println(next())
}
