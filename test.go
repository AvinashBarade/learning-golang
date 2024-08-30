package main

import (
	"fmt"
)

// twoSum function finds indices of the two numbers such that they add up to the target.
func twoSum(nums []int, target int) []int {
	numMaps := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if idx, found := numMaps[complement]; found {
			return []int{idx, i}
		}

		numMaps[num] = i

	}

	return []int{}
}

func reverse(num int) int {

	reverseNumber := 0
	originalNumber := num

	if num < 0 {
		return 0
	}

	for num > 0 {
		digit := num % 10
		reverseNumber = reverseNumber*10 + digit
		num = num / 10
	}

	fmt.Println(originalNumber, reverseNumber)
	if originalNumber == reverseNumber {

		return originalNumber
	} else {
		fmt.Println("Here")
		return 0
	}
}

func main() {
	// Example usage
	nums := []int{2, 7, 11, 15}
	target := 9

	// Call twoSum function
	result := twoSum(nums, target)

	// Print the result
	if len(result) == 2 {
		fmt.Printf("Indices: %d, %d\n", result[0], result[1])
	} else {
		fmt.Println("No two sum solution found.")
	}

	fmt.Println(reverse(121))
}
