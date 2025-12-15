package handlers

import (
    "fmt"
    "html/template"
    "net/http"
    "strconv"  > import string conversion pour convertir entre strings et nombres 
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        ErrorHandler(w, "Page non trouvée", 404)
        return
    }

    tmpl, err := template.ParseFiles("templates/home.html")
    if err != nil {
        ErrorHandler(w, "Erreur template", 500)
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
        ErrorHandler(w, "ID manquant", 400)
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        ErrorHandler(w, "ID invalide", 400)
        return
    }

    artist := GetTestArtistByID(id)
    if artist == nil {
        ErrorHandler(w, "Artiste non trouvé", 404)
        return
    }

    tmpl, err := template.ParseFiles("templates/artist-detail.html")
    if err != nil {
        ErrorHandler(w, "Erreur template", 500)
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





