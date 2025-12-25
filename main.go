package main

import (
	"fmt"
	"groupie-tracker/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artists", handlers.ArtistsHandler)
	http.HandleFunc("/artist", handlers.ArtistDetailHandler)

	http.HandleFunc("/test404", func(w http.ResponseWriter, r *http.Request) {
		handlers.ErrorHandler(w, 404, "Test page 404")
	})
	http.HandleFunc("/test400", func(w http.ResponseWriter, r *http.Request) {
		handlers.ErrorHandler(w, 400, "Test page 400")
	})
	http.HandleFunc("/test500", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" Route /test500 appelée ! Status = 500")
		handlers.ErrorHandler(w, 500, "Test page 500")
	})

	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
