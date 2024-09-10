package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://twitter.com",
		"http://golang.org",
		"http://github.com",
	}

	for _, link := range links {
		if checkStatus(link) {
			fmt.Printf("Status of %v is Live.\n", link)
		} else {
			fmt.Printf("Status of %v is DOWN!\n", link)
		}
	}

}

func checkStatus(site string) bool {
	resp, err := http.Get(site)

	if err != nil {
		return false
	}

	return resp.StatusCode == http.StatusOK
}
