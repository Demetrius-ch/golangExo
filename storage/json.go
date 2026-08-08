package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"mylibrary/models"
)

var (
	Livres     []models.Livre
	DernierID  = 0
	nomFichier = "livres.json"
)

// SauvegarderFichier écrit la slice Livres dans le fichier JSON
func SauvegarderFichier() error {
	donnees, err := json.MarshalIndent(Livres, "", "  ")
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON :", err)
		return err
	}
	return os.WriteFile(nomFichier, donnees, 0644)
}

// ChargerFichier lit le fichier JSON au démarrage ou initialise le stock par défaut
func ChargerFichier() {
	donnees, err := os.ReadFile(nomFichier)
	if err != nil {
		fmt.Println("Aucun fichier de sauvegarde trouvé. Initialisation du stock par défaut.")
		// Livres = []models.Livre{
		// 	{ID: 1, Titre: "Golang guide complet", Auteur: "CDEV Solutions", Disponible: false},
		// 	{ID: 2, Titre: "JavaScript", Auteur: "Believemy", Disponible: true},
		// }
		SauvegarderFichier()
	} else {
		err = json.Unmarshal(donnees, &Livres)
		if err != nil {
			fmt.Println("Erreur lors de la lecture des données JSON :", err)
			return
		}
	}

	maxID := 0
	for _, l := range Livres {
		if l.ID > maxID {
			maxID = l.ID
		}
	}
	DernierID = maxID
	fmt.Printf("Initialisation réussie : %d livre(s) chargé(s). Dernier ID = %d\n", len(Livres), DernierID)
}
