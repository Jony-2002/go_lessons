package main

import (
	"fmt"
	"html"
	"net/http"
)

const PORT = ":9999"

func goodMorning(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Good Morning</h1>")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func goodNight(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Body)
	// fmt.Println(r.Header)
	
	fmt.Println("Good Night...")
	fmt.Fprintf(w, "<h1>Good Night</h1>")
}

func main() {
	fmt.Println("Server is running on: 9999")

	// routes
	http.HandleFunc("/getMorning", goodMorning)
	http.HandleFunc("/getNight", goodNight)

	// listener
	http.ListenAndServe(PORT, nil)
}
