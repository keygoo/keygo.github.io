package hand

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)

	if r.URL.Path != "/about" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("page", "about.html"), path.Join("page", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Site Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":  "Retorica Puzzle | About",
		"banner": "Welcome to My Pages",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Site Error", http.StatusInternalServerError)
		return
	}
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)

	if r.URL.Path != "/contact" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("page", "contact.html"), path.Join("page", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Site Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":  "Retorica Puzzle | Contact",
		"banner": "Welcome to My Pages",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Site Error", http.StatusInternalServerError)
		return
	}
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)

	if r.URL.Path != "/blog" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("page", "blog.html"), path.Join("page", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Site Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":  "Retorica Puzzle | Blog",
		"banner": "Welcome to My Pages",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Site Error", http.StatusInternalServerError)
		return
	}
}
