package main

import "fmt"

func calculerImpot(salaireAnnuel int) int {
	impot := 0
	if salaireAnnuel > 3000000 {
		impot = ((salaireAnnuel - 3000000) / 5) + 200000
	} else if salaireAnnuel > 1000000 {
		impot = ((salaireAnnuel - 1000000) / 10)

	} else {
		impot = 0
	}
	return impot
}
func main() {
	monTaux := calculerImpot(35000000)
	fmt.Printf("Le taux de mon salaire annuel est %.2f", float64(monTaux)/100)
}
