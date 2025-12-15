package main

import (
	"html/template"
	"log"
	"net/http"
)

func StartServer() {

	artists, _ := SearchArtist()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, artists)

	})

	css := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	log.Println("Serveur en écoute sur https://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erreur au lancement du serveur: %v", err)
	}

}
