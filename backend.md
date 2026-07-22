package main

import (
"encoding/json"
"fmt"
"net/http"
)

type Produit struct {
ID int `json:"id"`
Nom string `json:"nom"`
Prix float64 `json:"prix"`
EnStock bool `json:"en_stock"`
}

func handler(w http.ResponseWriter, r *http.Request) {
fmt.Fprintln(w, "Bienvenue sur mon serveur Golang")
}
func produitHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

    p := []Produit{
    	{
    		ID:      1,
    		Nom:     "Ecran 4K",
    		Prix:    299.99,
    		EnStock: true,
    	},
    	{
    		ID:      2,
    		Nom:     "Clavier Mécanique",
    		Prix:    79.5,
    		EnStock: false,
    	},
    	{
    		ID:      3,
    		Nom:     "Souris Sans fil",
    		Prix:    35,
    		EnStock: true,
    	},
    }

    json.NewEncoder(w).Encode(p)

}

func main() {
http.HandleFunc("/", handler)
http.HandleFunc("/api/produits", produitHandler)
fmt.Println("Le serveur démarre sur le port http://localhost:8080")

    err := http.ListenAndServe(":8080", nil)

    if err != nil {
    	fmt.Println("Erreur", err)
    }

}
omment lire du JSON envoyé par le client ?
Quand un client (React, Flutter, Postman, Curl) t'envoie un formulaire ou un produit, la donnée voyage dans r.Body (le corps de la requête).

Pour la récupérer, la démarche se fait en 4 étapes :

Vérifier la méthode HTTP : On veut s'assurer que c'est bien une requête POST (et pas un GET).

Préparer une variable vide du type de notre struct.

Décoder le JSON de r.Body dans cette variable avec json.NewDecoder(r.Body).Decode(&produit).

Valider et traiter la donnée.

package main

import (
"encoding/json"
"fmt"
"net/http"
"strconv"
)

type Task struct {
ID int `json:"id"`
Titre string `json:"titre"`
Terminee bool `json:"terminee"`
}

var tasks = []Task{
{
ID: 1,
Titre: "",
Terminee: false,
},
{
ID: 2,
Titre: "Créer une API",
Terminee: true,
},
}

func handler(w http.ResponseWriter, r \*http.Request) {
fmt.Fprintln(w, "Bienvenue sur mon serveur")
}

func tasksHandler(w http.ResponseWriter, r \*http.Request) {
w.Header().Set("Content-Type", "application/json")

    if r.Method == http.MethodGet {
    	idStr := r.URL.Query().Get("id")
    	if idStr == "" {
    		json.NewEncoder(w).Encode(tasks)
    		return
    	}
    	if idStr != "" {
    		idRecherche, err := strconv.Atoi(idStr)
    		if err != nil {
    			w.WriteHeader(http.StatusBadRequest)
    			fmt.Fprintln(w, ` "{Erreur l'ID invalide}"`)
    			return
    		}
    		for _, task := range tasks {
    			if task.ID == idRecherche {
    				json.NewEncoder(w).Encode(task)
    				return
    			}
    			// if task.ID != idRecherche {
    			// 	w.WriteHeader(http.StatusBadRequest)
    			// 	fmt.Fprintln(w, ` "{Erreur l'ID  recherchée invalide}"`)
    			// 	return

    			// }
    		}
    		w.WriteHeader(http.StatusNotFound)
    		fmt.Fprintln(w, `{"erreur": "Tâche non trouvée"}`)
    		return
    	}
    }
    if r.Method == http.MethodPost {
    	var nouvelleTask Task
    	err := json.NewDecoder(r.Body).Decode(&nouvelleTask)

    	if err != nil {
    		w.WriteHeader(http.StatusBadRequest)
    		fmt.Fprintln(w, "Erreur")
    		return
    	}
    	if nouvelleTask.Titre == "" {
    		w.WriteHeader(http.StatusBadRequest)
    		fmt.Fprintln(w, "Le champ titre est vide. ")
    		return
    	}

    	nouvelleTask.ID = len(tasks) + 1
    	tasks = append(tasks, nouvelleTask)
    	w.WriteHeader(http.StatusCreated)

    	json.NewEncoder(w).Encode(nouvelleTask)
    	return

    }
    w.WriteHeader(http.StatusMethodNotAllowed)

}

func main() {
http.HandleFunc("/", handler)
http.HandleFunc("/api/tasks", tasksHandler)

    fmt.Println("Le serveur démarre sur le port http://localhost:8080")
    http.ListenAndServe(":8080", nil)

}
