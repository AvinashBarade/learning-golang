package main

import (
	"fmt"
	"net/http"
	"sync"
)

func checkLink(link string, wg *sync.WaitGroup) {

	defer wg.Done()
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link + " is down")
	}
	fmt.Println(link + " is up")

}

func main() {
	var wg sync.WaitGroup
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		wg.Add(1)
		go checkLink(link, &wg)
	}
	wg.Wait()
}
