package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)
type ErrorData struct {
    Status  int
    Message string
}
func showError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	data := ErrorData{
		Status:  status,
		Message: message,
	}

	tmpl.Execute(w, data)
}

func _Handler() {

	artists, err := SearchArtist()
	if err != nil {
		log.Fatalf("Erreur SearchArtist: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			showError(w, 404, "")
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		err = tmpl.Execute(w, artists)
		if err != nil {
			showError(w, 500, "Erreur lors de l'affichage de la page")
		}
	})

	http.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/map.html"))
		err = tmpl.Execute(w, artists)
		if err != nil {
			showError(w, 500, "Erreur lors de l'affichage de la carte")
		}
	})

	http.HandleFunc("/artistes", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			showError(w, 400, "ID d'artiste manquant")
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			showError(w, 400, "ID d'artiste invalide")
			return
		}

		if id <= 0 {
			showError(w, 400, "ID doit être supérieur à 0")
			return
		}

		if id > 52 {
			showError(w, 404, "Artiste introuvable")
			return
		}

		artist, err := getOneArtist(id)
		if err != nil {
			showError(w, 500, "Erreur lors de la récupération de l'artiste")
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/artistes.html"))
		err = tmpl.Execute(w, artist)
		if err != nil {
			showError(w, 500, "Erreur lors de l'affichage de la page artiste")
		}
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/contact.html"))
		err = tmpl.Execute(w, artists)
		if err != nil {
			showError(w, 500, "Erreur lors de l'affichage de la page de contact")
		}
	})
}
