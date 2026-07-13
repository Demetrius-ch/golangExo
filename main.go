package main

import "fmt"

type Smartphone struct {
	Modele   string
	Batterie int
}

// Correction 1 : Nom de la méthode corrigé ("charger") et récepteur clair ("s")
func (s *Smartphone) charger(minutes int) {
	gain := minutes * 2
	s.Batterie += gain

	if s.Batterie > 100 {
		s.Batterie = 100
		fmt.Println("Charge complète atteinte !")
	}
}

func main() {
	monTel := Smartphone{
		Modele:   "Tecno",
		Batterie: 20,
	}

	// Correction 2 : Affichage AVANT toute action
	fmt.Printf("Avant charge : %s à %d%%\n", monTel.Modele, monTel.Batterie)

	// Première charge (10 min = +20%) -> Devient 40%
	monTel.charger(10)

	// Deuxième charge (50 min = +100% théorique) -> Capé à 100%
	monTel.charger(50)

	fmt.Printf("Après charge : %s à %d%%\n", monTel.Modele, monTel.Batterie)
}   