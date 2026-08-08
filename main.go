package main

import (
	"fmt"
	"net/http"

	"mylibrary/handlers"
	"mylibrary/storage"
)

func main() {
	// Charger les données au démarrage
	storage.ChargerFichier()

	// Enregistrement des routes
	http.HandleFunc("/api/livres", handlers.LoggerMiddleware(handlers.HandlerLivre))

	fmt.Println("Le serveur démarre sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
