package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, 404, "Page non trouvée")
		return
	}

	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Erreur template", 500)
		return
	}

	tmpl.Execute(w, nil)
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists := GetTestArtists()

	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		ErrorHandler(w, 500, "Erreur template")
		fmt.Println("Erreur template:", err)
		return
	}

	err = tmpl.Execute(w, artists)
	if err != nil {
		fmt.Println("Erreur exécution:", err)
	}
}

func ArtistDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		ErrorHandler(w, 400, "ID manquant")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorHandler(w, 400, "ID invalide")
		return
	}

	artist := GetTestArtistByID(id)
	if artist == nil {
		ErrorHandler(w, 404, "Artiste non trouvé")
		return
	}

	tmpl, err := template.ParseFiles("templates/artist-detail.html")
	if err != nil {
		ErrorHandler(w, 500, "Erreur template")
		return
	}

	tmpl.Execute(w, artist)
}

func ErrorHandler(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, message, status)
		return
	}

	data := map[string]interface{}{
		"Status":  status,
		"Message": message,
	}

	tmpl.Execute(w, data)
}
