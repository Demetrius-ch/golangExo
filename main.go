package main

import (
	"fmt"
	"net/http"
)

func acceuilHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenue sur mon serveur golang")
}
func monHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Demetrius, ton serveur go t'envoie cette réponse.")
}
func saluerHandler(w http.ResponseWriter, r *http.Request) {
	nom := r.URL.Query().Get("nom")
	if nom == "" {
		nom = "Visiteur"
	}
	fmt.Fprintf(w, "Bonjour %s, ravi de te voir sur l'API !", nom)

}

func main() {
	http.HandleFunc("/", acceuilHandler)
	http.HandleFunc("/hello", monHandler)
	http.HandleFunc("/saluer", saluerHandler)
	port := "http://localhost:8080"

	fmt.Printf("Le serveur démarre sur le port %v\n", port)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
