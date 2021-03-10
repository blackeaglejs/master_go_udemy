package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		c <- s //sending into the channel
	} else {
		defer resp.Body.Close() // Prevents memory leaks and makes sure that any  blocks are removed.
		s := fmt.Sprintf("Status Code: %d  \n", resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"
			s += fmt.Sprintf("Writing response body to %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				// logger.Fatal(err)
				s += "Error writing file\n"
				c <- s
			}
		}
	}
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	c := make(chan string)

	for _, url := range urls {
		go checkAndSaveBody(url, c)
		fmt.Println(strings.Repeat("#", 20))
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

	// This version of this program is not concurrent and will take nearly 34 minutes to execute.
}
