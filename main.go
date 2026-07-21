package main

import (
    "fmt"
    "os"
)

func main() {
    // On lit le fichier. Ça nous renvoie des bytes et une erreur
    donnees, err := os.ReadFile("notes.txt")
    if err != nil {
        fmt.Println("Erreur lors de la lecture :", err)
        return
    }
    
    // On transforme les bytes en texte lisible (string)
    fmt.Println(string(donnees))
}