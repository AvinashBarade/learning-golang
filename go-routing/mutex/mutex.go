package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	counter = 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
		fmt.Println(counter)
	}
	wg.Wait()
	fmt.Println("Final:", counter)
}
