package main

import (
	"log"
	"net/http"

	"github.com/retorica-puzzle/retorica-puzzle.github.io/hand"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hand.HomeHandler)
	mux.HandleFunc("/about", hand.AboutHandler)
	mux.HandleFunc("/contact", hand.ContactHandler)
	mux.HandleFunc("/blog", hand.BlogHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	log.Println("Start Web Server")

	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
