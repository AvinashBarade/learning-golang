package main

import "fmt"

type Location struct {
	X, Y float32
}

func (l Location) printLocation() {
	fmt.Println(l.X, l.Y)
}

func main() {
	pune := Location{2.45, 6.78}
	pune.printLocation()
}
