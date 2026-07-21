package main

import (
	"fmt"
	"os"
)

func main() {
	texte := "Bonjour Demetrius, ceci est enregistré sur votre disque."

	donnees := []byte(texte)

	err := os.WriteFile("notes.txt", donnees, 0644)

	if err != nil {
		//	Gérer les permissions même en cas d'erreur
	}
	_, errf := os.ReadFile("note.txt")

	if errf != nil {
		fmt.Println("Erreur lors de la lecteur du fichier", errf)
		return
	}
	fmt.Println(string(donnees))

}
