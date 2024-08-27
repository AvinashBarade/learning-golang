package main

import "fmt"

func main() {
	n := 0

	fmt.Scanf("%d", &n)
	singerCount := make(map[int]int)

	for i := 0; i < n; i++ {
		var singer int
		fmt.Scan(&singer)
		singerCount[singer]++
	}

	maxCount := 0
	for _, count := range singerCount {
		if count > maxCount {
			maxCount = count
		}
	}

	noSiger := 0
	for _, count := range singerCount {
		if count == maxCount {
			noSiger++
		}
	}

	fmt.Println(noSiger)

}
