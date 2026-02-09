package main

import (
	"log"
	"net/http"
)

//StartServer starts the HTTP server on port 8080
// Configures routes via _Handler() and serves static files (CSS/JS)
// Stops the program in case of a startup error
func StartServer() {
	// Routes HTML
	_Handler()

	// Static files
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	log.Println("Serveur en écoute sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
