package fact

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(fact(num))
}

func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num-1)
}
