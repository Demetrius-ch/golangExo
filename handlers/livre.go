package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"mylibrary/models"
	"mylibrary/storage"
)

func HandlerGetLivres(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		json.NewEncoder(w).Encode(storage.Livres)
		return
	}

	idRecherche, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"erreur": "L'ID doit être un nombre valide"}`)
		return
	}

	for _, livre := range storage.Livres {
		if livre.ID == idRecherche {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(livre)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, `{"erreur": "Livre non trouvé"}`)
}

func HandlerCreateLivre(w http.ResponseWriter, r *http.Request) {
	var nouveauLivre models.Livre

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

	storage.DernierID++
	nouveauLivre.ID = storage.DernierID
	nouveauLivre.Disponible = true

	storage.Livres = append(storage.Livres, nouveauLivre)
	storage.SauvegarderFichier()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nouveauLivre)
}

func HandlerUpdateLivre(w http.ResponseWriter, r *http.Request) {
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

	var nouveauLivre models.Livre
	err = json.NewDecoder(r.Body).Decode(&nouveauLivre)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"erreur": "JSON invalide"}`)
		return
	}

	for i, livre := range storage.Livres {
		if idRecherche == livre.ID {
			storage.Livres[i].Titre = nouveauLivre.Titre
			storage.Livres[i].Auteur = nouveauLivre.Auteur
			storage.Livres[i].Disponible = nouveauLivre.Disponible
			storage.SauvegarderFichier()

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(storage.Livres[i])
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, `{"erreur": "Livre introuvable"}`)
}

func HandlerDeleteLivre(w http.ResponseWriter, r *http.Request) {
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

	for i, livre := range storage.Livres {
		if idRecherche == livre.ID {
			livreSupprime := storage.Livres[i]
			storage.Livres = append(storage.Livres[:i], storage.Livres[i+1:]...)
			storage.SauvegarderFichier()

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(livreSupprime)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, `{"erreur": "Livre non trouvé"}`)
}

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s\n", r.Method, r.URL.Path)
		next(w, r)
	}
}

func HandlerLivre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		HandlerGetLivres(w, r)
	case http.MethodPost:
		HandlerCreateLivre(w, r)
	case http.MethodDelete:
		HandlerDeleteLivre(w, r)
	case http.MethodPut:
		HandlerUpdateLivre(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
