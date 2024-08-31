package main

import "fmt"

func main() {
	//var num int
	//fmt.Scan(&num)
	fmt.Println(fibo(5))
}

func fibo(num int) int {
	if num == 0 || num == 1 {
		return num
	}

	return fibo(num-1) + fibo(num-2)
}
