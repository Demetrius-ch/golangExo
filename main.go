package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var dernierArticle int

func init() {
	maxID := 0
	for _, article := range articles {
		if article.ID > maxID {
			maxID = article.ID
		}
	}
	dernierArticle = maxID
	fmt.Printf("Initialisation : Dernier ID détecté = %d\n", dernierArticle)
}

type Article struct {
	ID       int    `json:"id"`
	Nom      string `json:"nom"`
	Quantite int    `json:"quantite"`
}

var articles = []Article{
	{
		ID:       1,
		Nom:      "Pot de peinture",
		Quantite: 20,
	},
	{
		ID:       2,
		Nom:      "Crème",
		Quantite: 5,
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenue sur mon Serveur Golang")
}

func handleGetArticles(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		json.NewEncoder(w).Encode(articles)
		return
	}
	if idStr != "" {
		idRecherche, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Id non trouvée")
			return
		}
		for _, article := range articles {
			if article.ID == idRecherche {
				json.NewEncoder(w).Encode(article)
				return
			}

		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, `"{l'ID n'existe pas}"`)
	}

}

func handleCreateArticles(w http.ResponseWriter, r *http.Request) {
	var nouvelArticle Article
	err := json.NewDecoder(r.Body).Decode(&nouvelArticle)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `"{Erreur lors de la création}"`)
		return
	}
	if nouvelArticle.Nom == "" || nouvelArticle.Quantite < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `"{Le champ Article doit être obligatoire et Quantité doit être strictement suprérieur à 0}"`)
		return

	}
	dernierArticle++
	//nouvelArticle.ID = len(articles) + 1
	nouvelArticle.ID = dernierArticle
	articles = append(articles, nouvelArticle)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nouvelArticle)
	//	return

}

func handlerDeleteArticle(w http.ResponseWriter, r *http.Request) {
	// var nouvelleArticle Article
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"erreur": "ID requis pour la suppression"}`)
		return
	}
	if idStr != "" {
		idRecherche, err := strconv.Atoi(idStr)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, `{"erreur": "ID requis pour la suppression"}`)
			return

		}
		for i, article := range articles {
			if article.ID == idRecherche {

				articles = append(articles[:i], articles[i+1:]...)
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, `{"message": "Article supprime"}`)
				json.NewEncoder(w).Encode(article)
				return
			}

		}

		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, `{"erreur": "Article non trouvé"}`)

	}
	//nouvelleArticle.ID = len(articles) - 1
	//articles = append(articles, nouvelleArticle)
	//w.WriteHeader(http.StatusLoopDetected)
	//json.NewEncoder(w).Encode(nouvelleArticle)

}
func handlerUpdateArticle(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	var nouvelArticle Article

	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Le champ id doit être obligatoire")
		return
	}
	idRecherche, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `"{Erreur" : "ID invalide"`)
		return

	}
	err = json.NewDecoder(r.Body).Decode(&nouvelArticle)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `"{Erreur" : "JSON invalide"`)
		return

	}
	if nouvelArticle.Nom == "" || nouvelArticle.Quantite < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"Erreur": "Nom obligatoire et Quantité >=0"}`)
	}

	for i, article := range articles {
		//idRecherche, _ := strconv.Atoi(idStr)
		if article.ID == idRecherche {
			articles[i].Nom = nouvelArticle.Nom
			articles[i].Quantite = nouvelArticle.Quantite
			nouvelArticle.ID = idRecherche
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Article a été mise à jour avec succès.")
			json.NewEncoder(w).Encode(articles[i])

			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "ID introuvable")

	}

}

// Middleware de logging
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Affiche : [METHODE] URL (Ex: [GET] /api/articles)
		fmt.Printf("[%s] %s\n", r.Method, r.URL.Path)
		next(w, r) // Passe la main au handler normal
	}
}

func handlerArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		handleGetArticles(w, r)
	case http.MethodPost:
		handleCreateArticles(w, r)
	case http.MethodDelete:
		handlerDeleteArticle(w, r)
	case http.MethodPut:
		handlerUpdateArticle(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/api/articles", handlerArticle)
	fmt.Println("Le serveur démarre sur le port http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
