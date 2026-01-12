package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func _Handler() {

	artists, err := SearchArtist()
	if err != nil {
		log.Fatalf("Erreur SearchArtist: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		_ = tmpl.Execute(w, artists)
	})

	http.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/map.html"))
		_ = tmpl.Execute(w, artists)
	})

	http.HandleFunc("/artistes", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		artist, _ := getOneArtist(id)
		tmpl := template.Must(template.ParseFiles("templates/artistes.html"))
		_ = tmpl.Execute(w, artist)
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/contact.html"))
		_ = tmpl.Execute(w, artists)
	})
}
