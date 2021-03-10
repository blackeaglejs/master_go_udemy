package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func checkURL(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v  \n", err)
		fmt.Println(s)
		c <- url
	} else {
		s := fmt.Sprintf("Status Code: %d  \n", resp.StatusCode)
		s += fmt.Sprintf("%s is UP\n", url)
		fmt.Println(s)
		c <- url
	}
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	c := make(chan string)

	for _, url := range urls {
		go checkURL(url, c)
		fmt.Println(strings.Repeat("#", 20))
	}

	// for {
	// 	go checkURL(<-c, c)
	// 	fmt.Println(strings.Repeat("#", 30))
	// 	time.Sleep(time.Second * 2)
	// }

	// Alternate syntax using range.
	for url := range c {
		go func(u string) {
			time.Sleep(time.Second * 2)
			go checkURL(u, c)
		}(url)
	}

	// This version of this program is not concurrent and will take nearly 34 minutes to execute.
}
