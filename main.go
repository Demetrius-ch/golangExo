package main

import "fmt"

// 1. On définit le contrat
type Chargeable interface {
	charger(minutes int)
}

// 2. Notre Smartphone (ton code !)
type Smartphone struct {
	Modele   string
	Batterie int
}
func (s *Smartphone) charger(minutes int) {
	s.Batterie += minutes * 2
	if s.Batterie > 100 { s.Batterie = 100 }
}

// 3. Une nouvelle structure : la Voiture Électrique
type VoitureElectrique struct {
	Modele string
	Autonomie int
}
func (v *VoitureElectrique) charger(minutes int) {
	v.Autonomie += minutes * 5 // Charge plus vite !
	if v.Autonomie > 100 { v.Autonomie = 100 }
}

// 4. LA FONCTION MAGIQUE
// Elle ne prend ni un téléphone, ni une voiture. Elle prend n'importe quoi qui est "Chargeable"
type BorneDeRecharge struct{}

func (b BorneDeRecharge) Alimenter(appareil Chargeable, minutes int) {
	fmt.Println("Connexion établie... Début de la charge.")
	appareil.charger(minutes) // Go sait quelle méthode appeler selon l'objet !
}

func main() {
	tel := Smartphone{Modele: "Tecno", Batterie: 20}
	tesla := VoitureElectrique{Modele: "Model 3", Autonomie: 10}

	borne := BorneDeRecharge{}

	// La borne peut charger le téléphone...
	borne.Alimenter(&tel, 20)
	
	// ... et la même borne peut charger la Tesla !
	borne.Alimenter(&tesla, 10)

	fmt.Println("Batterie Tel :", tel.Batterie)      // 60%
	fmt.Println("Autonomie Tesla :", tesla.Autonomie) // 60%
}