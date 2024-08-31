package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	str1 := "go lang in simple and fast"
	str2 := "fast"
	fmt.Println(strings.Contains(str1, str2))
	fmt.Println(runtime.NumCPU())

}
