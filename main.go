package main

import (
	"fmt"
	"myapp/pages"
	"net/http"
)

const portNumber = ":8080"

// web server is built-in to standard library
func main() {
	http.HandleFunc("/", pages.HomePage)
	http.HandleFunc("/about", pages.AboutPage)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
