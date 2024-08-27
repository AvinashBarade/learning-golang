package main

import "fmt"

func main() {
	i := 10
	k := &i
	fmt.Println(*k)
	*k = 20
	fmt.Println(i)
	fmt.Println(*k)
}
