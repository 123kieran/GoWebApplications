// Problem 1
// Kieran O'Haloran 25/10/17

package main

import (
	"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	//Browser renders html tags
	w.Header().Set("Content-Type", "text/html")

	//Output to browser
	fmt.Fprintln(w, "Guessing Game")
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}