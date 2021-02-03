package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func compile(w http.ResponseWriter, r *http.Request) {
	requestString := r.URL.Path[1:]
	lex := lex(requestString)
	a, b := parse(lex)
	returnString := transpile(a, b)
	fmt.Fprintf(w, returnString)
}

func setupRoutes() {
	http.HandleFunc("/", compile)
}
