// Kieran O'Haloran 25/10/17
//http://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html
//// https://data-representation.github.io

package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type myMsg struct {
	Message   string
	YourGuess int
}

func server(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	http.ServeFile(w, r, "index.html")
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	message := "Guess a number between 1 and 20"
	rand.Seed(time.Now().UTC().UnixNano())

	target := 0
	var cookie, err = r.Cookie("target")

	if err == nil {

		target, _ = strconv.Atoi(cookie.Value)
		if target == 0 {
			target = rand.Intn(20 - 1)
		}
	}

	yourGuess, _ := strconv.Atoi(r.FormValue("guess"))
	//compare YourGuess to target guess(random number)
	if yourGuess == target {
		message = "Correct Guess " + strconv.Itoa(yourGuess) + " was the answer"
	} else if yourGuess < target {
		message = "Your Guess is too low, Try again!"
	} else {
		message = "Your Guess is too high, Try again!"
	}

	cookie = &http.Cookie{
		Name:    "target",
		Value:   strconv.Itoa(target),
		Expires: time.Now().Add(72 * time.Hour),
	}

	http.SetCookie(w, cookie)

	t, _ := template.ParseFiles("guess.tmpl")

	t.Execute(w, &myMsg{Message: message})
}

func main() {
	http.HandleFunc("/", server) //"/"handles any requests and passes to server
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil) //webserver is started
}
