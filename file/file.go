package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.go")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 1000)

	n, err := file.Read(buf)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Read %d bytes: %s\n", n, string(buf[:n]))
}
