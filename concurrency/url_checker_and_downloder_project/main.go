package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"bitbucket.org/accelecon/charybdis/services/device-grpc-api/logger"
)

func checkAndSaveBody(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is DOWN!\n", url)
	} else {
		defer resp.Body.Close() // Prevents memory leaks and makes sure that any  blocks are removed.
		fmt.Printf("%s => Status Code: %d \n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"
			fmt.Printf("Writing response body to %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				logger.Fatal(err)
			}
		}
	}
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	for _, url := range urls {
		checkAndSaveBody(url)
		fmt.Println(strings.Repeat("#", 20))
	}

	// This version of this program is not concurrent and will take nearly 34 minutes to execute.
}
