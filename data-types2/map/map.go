package main

import "fmt"

func main() {

	//var colors map[string]string

	colors := make(map[string]string)
	colors["a"] = "avinash"
	colors["b"] = "bavashir"
	colors["c"] = "She"

	fmt.Println(colors)
	delete(colors, "c")
	fmt.Println(colors)

	interatingMap()
}

func interatingMap() {
	colors := map[string]string{
		"red":   "ajdbhdb",
		"black": "auhbyb",
		"a":     "a",
	}

	for i, v := range colors {
		fmt.Printf("%s = %s ", i, v)
	}
}
