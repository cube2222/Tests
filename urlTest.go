package main

import (
	"net/http"
	"fmt"
)

func main() {
	r, _ := http.NewRequest("GET", "https://meetuprest.appspot.com/public/html/presentation.html?key=123123123", nil)
	fmt.Printf("URL: %v\n", r.URL)
	fmt.Printf("Request url: %v\n", r.URL.RequestURI())
	fmt.Printf("Raw path: %v\n", r.URL.RawPath)
	fmt.Printf("Raw query: %v\n", r.URL.RawQuery)
	fmt.Printf("Fragment %v\n", r.URL.Fragment)
}
