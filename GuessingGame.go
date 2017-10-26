// Problem 2
// Kieran O'Haloran 25/10/17

package main

import (
	"fmt"
	"net/http"
)



func server(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "<h1>Guessing Game</h1>")
	http.ServeFile(w, r, "guess.html")}


func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
