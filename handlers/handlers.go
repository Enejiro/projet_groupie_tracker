package handlers

import (
    "html/template"
    "net/http"
    "projet_groupie_tracker/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/home.html")
    if err != nil {
        http.Error(w, "Erreur template", 500)
        return
    }

    artist := models.Artist{
        Name:         "Queen",
        CreationDate: 1970,
        Members:      4,
    }

    tmpl.Execute(w, artist)
}
