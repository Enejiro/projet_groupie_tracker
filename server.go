package main

import (
    "fmt"
    "log"
    "net/http"
)

func StartServer() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, world!")
    })

    log.Println("Serveur en écoute sur http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Erreur au lancement du serveur: %v", err)
    }
}
