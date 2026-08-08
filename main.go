package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Livre struct {
	ID         int    `json:"id"`
	Titre      string `json:"titre"`
	Auteur     string `json:"auteur"`
	Disponible bool   `json:"disponible"`
}

var livres = []Livre{
	{
		ID:         1,
		Titre:      "Golang guide complet",
		Auteur:     "CDEV Solutions",
		Disponible: false,
	},
	{
		ID:         2,
		Titre:      "JavaScript",
		Auteur:     "Believemy",
		Disponible: true,
	},
}

var dernierID = 0

func init() {
	var maxId = 0
	for _, l := range livres {
		if l.ID > maxId {
			maxId = l.ID // Fix 1 : Mise à jour de maxId
		}
	}
	dernierID = maxId
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenue sur mon serveur Golang")
}

func handlerGetLivres(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		json.NewEncoder(w).Encode(livres)
		return
	}

	idRecherche, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"erreur": "L'ID doit être un nombre valide"}`) // Fix 2 : Fprintln avec w
		return
	}

	for _, livre := range livres {
		if livre.ID == idRecherche {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(livre)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound) // Fix 6 : 404 Not Found
	fmt.Fprintln(w, `{"erreur": "Livre non trouvé"}`)
}

func handlerCreateLivre(w http.ResponseWriter, r *http.Request) {
	var nouveauLivre Livre

	err := json.NewDecoder(r.Body).Decode(&nouveauLivre)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"erreur": "Erreur lors de la création"}`)
		return
	}

	if nouveauLivre.Titre == "" || nouveauLivre.Auteur == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"erreur": "Les champs titre et auteur sont obligatoires"}`)
		return
	}

	dernierID++
	nouveauLivre.ID = dernierID // Fix 3 : Assignation de l'ID généré
	nouveauLivre.Disponible = true

	livres = append(livres, nouveauLivre)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nouveauLivre) // Fix 3 : Renvoie uniquement le livre créé
}

func handlerUpdateLivre(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"erreur": "L'ID est obligatoire dans l'URL"}`)
		return
	}

	idRecherche, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"erreur": "ID invalide"}`)
		return
	}

	var nouveauLivre Livre
	err = json.NewDecoder(r.Body).Decode(&nouveauLivre)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"erreur": "JSON invalide"}`)
		return
	}

	for i, livre := range livres {
		if idRecherche == livre.ID {
			// Fix 4 : Mis à jour avec les champs du JSON reçu
			livres[i].Titre = nouveauLivre.Titre
			livres[i].Auteur = nouveauLivre.Auteur
			livres[i].Disponible = nouveauLivre.Disponible

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(livres[i])
			return
		}
	}

	// Fix 4 : Message d'erreur en dehors de la boucle
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, `{"erreur": "Livre introuvable"}`)
}

func handlerDeleteLivre(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"erreur": "L'ID est obligatoire dans l'URL"}`)
		return
	}

	idRecherche, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"erreur": "ID invalide"}`)
		return
	}

	for i, livre := range livres {
		if idRecherche == livre.ID {
			livreSupprime := livres[i]
			// Fix 5 : Slicing correct pour supprimer l'élément
			livres = append(livres[:i], livres[i+1:]...)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(livreSupprime)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, `{"erreur": "Livre non trouvé"}`)
}

// Middleware de logging
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s\n", r.Method, r.URL.Path)
		next(w, r)
	}
}

func handlerLivre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		handlerGetLivres(w, r)
	case http.MethodPost:
		handlerCreateLivre(w, r)
	case http.MethodDelete:
		handlerDeleteLivre(w, r)
	case http.MethodPut:
		handlerUpdateLivre(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	// Fix 7 : Enrobage avec loggerMiddleware
	http.HandleFunc("/", loggerMiddleware(handler))
	http.HandleFunc("/api/livres", loggerMiddleware(handlerLivre))

	fmt.Println("Le serveur démarre sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
