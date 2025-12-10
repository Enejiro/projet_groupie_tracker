package main

import (
    "log"
    "net/http"
)

func StartServer() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "templates/index.html")
    })

    log.Println("Serveur en écoute sur http://localhost:8888")
    if err := http.ListenAndServe(":8888", nil); err != nil {
        log.Fatalf("Erreur au lancement du serveur: %v", err)
    }
}

