package main

import (
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/retorica-puzzle/retorica-puzzle.github.io/hand"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/about", hand.AboutHandler)
	mux.HandleFunc("/contact", hand.ContactHandler)
	mux.HandleFunc("/blog", hand.BlogHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// log.Println("Start Web Server")

	// http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("index.html"), path.Join("page", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Site Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":  "My Home",
		"banner": "Welcome to My Pages",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Site Error", http.StatusInternalServerError)
		return
	}

}
