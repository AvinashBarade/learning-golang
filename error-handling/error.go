package main

import (
	"errors"
	"fmt"
)

func divide(A, B float64) (float64, error) {
	if B == 0 {
		return 0, errors.New("Cannot devide by Zero")
	}

	return A / B, nil

}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Result:", result)
}
