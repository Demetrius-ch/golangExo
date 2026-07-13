package main

import "fmt"

type Smartphone struct {
	Modele   string
	Batterie int
}

func (charger *Smartphone) chagerTelephone(minutes int) {
	gain := minutes * 2
	charger.Batterie += gain

	if charger.Batterie > 100 {
		charger.Batterie = 100
		fmt.Printf("Charge complète")
	}
}

func main() {
	monTel := Smartphone{
		Modele:   "Tecno",
		Batterie: 20,
	}
	monTel.chagerTelephone(10)

	fmt.Printf("Avant charge : %s à %d%%\n", monTel.Modele, monTel.Batterie)

	// Appel de la méthode : charge de 50 minutes (gain théorique de 100%)
	monTel.chagerTelephone(50)

	fmt.Printf(" Après charge : %s à %d%%\n", monTel.Modele, monTel.Batterie)
}
