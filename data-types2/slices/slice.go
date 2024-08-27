package main

import (
	"fmt"
	"strings"
)

func main() {
	primes := [6]int{1, 3, 4, 5, 6, 7}
	var s []int = primes[2:3]
	fmt.Println(s)
	sliceLiterals()
	sliceDefault()
	capacity()
	nilSlice()
	makeSlice()
	multiSlice()
	rangeSlice()
}

func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func sliceDefault() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[0:6]
	fmt.Println(s)
	fmt.Println(s[:])
	fmt.Println(s[3:])
	// s = s[:2]
	// fmt.Println(s)

	// s = s[1:]
	// fmt.Println(s)
	// s = s[:]
	// fmt.Println(s)
}

func capacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func nilSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

}

func makeSlice() {
	a := make([]int, 5)
	printSlice1("a", a)

	b := make([]int, 0, 5)
	printSlice1("b", b)

	c := b[:2]
	printSlice1("c", c)

	d := c[2:5]
	printSlice1("d", d)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func multiSlice() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][1] = "0"
	board[1][1] = "X"
	board[2][2] = "P"
	//fmt.Println(board)
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], ""))
	}
}

func appendslice() {
	var s []int
	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func rangeSlice() {
	var s = []int{4, 2, 4, 56, 7, 87}
	for _, v := range s {
		fmt.Printf("%d\n", v)
	}
}
