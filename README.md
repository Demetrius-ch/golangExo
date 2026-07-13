# Mes Exercices Go

## Table des matières

1. [Exercice 1 : Le Filtre Élite](#exercice-1--le-filtre-élite)
2. [Exercice 2 : Le Calculateur de Paie](#exercice-2--le-calculateur-de-paie)
3. [Cas Pratique : Système de Facturation Pay-As-You-Go](#cas-pratique--système-de-facturation-pay-as-you-go)
4. [Concepts : Slices et Tableaux](#concepts--slices-et-tableaux)
5. [Étape 5 : Les Maps (Dictionnaires Clés/Valeurs)](#étape-5--les-maps-dictionnaires-clésvaleurs)
6. [À Retenir](#à-retenir)

---

## Exercice 1 : Le Filtre Élite

### Objectif

Créer un système de filtrage de scores avec conditions.

### Énoncé

Écris une fonction nommée `analyserScore` qui prend un score `int` en paramètre et ne retourne rien.

**Règles :**

- Si le score est ≥ 80 : affiche "Score [X] : Reçu avec mention Élite !"
- Si le score est entre 50 et 79 : affiche "Score [X] : Reçu."
- Si le score est < 50 : affiche "Score [X] : Échoué."

**Dans ton main :**
Fais une boucle for qui teste des scores de 40 à 90, en incrémentant de 5 en 5 (40, 45, 50, 55...).

### Correction

```go
package main

import "fmt"

func analyserScore(score int) {
	if score >= 80 {
		fmt.Printf("Score %d : Reçu avec mention Elite !\n", score)
	} else if score >= 50 {
		fmt.Printf("Score %d : Reçu.\n", score)
	} else {
		fmt.Printf("Score %d : Échoué.\n", score)
	}
}

func main() {
	for score := 40; score <= 90; score += 5 {
		analyserScore(score)
	}
}
```

---

## Exercice 2 : Le Calculateur de Paie

### Objectif

Calculer le salaire d'un employé avec gestion des heures supplémentaires (retours multiples).

### Énoncé

Écris une fonction nommée `calculerSalaire` qui prend deux paramètres : `heuresTravaillees int` et `tauxHoraire int`.

**Cette fonction doit retourner deux valeurs (int, int) :**

- Le salaire de base (les heures jusqu'à 35 heures payées au taux normal)
- Le salaire majoré (les heures au-delà de 35 heures, payées le double du taux normal)

**Dans ton main :**

- Appelle cette fonction pour un employé qui a travaillé 42 heures à un taux de 20€/heure
- Récupère les deux résultats
- Affiche le bilan : "Salaire de base : X€, Heures sup : Y€, Total : Z€"

### Correction

```go
package main

import "fmt"

func calculerSalaire(heuresTravaillees, tauxHoraire int) (int, int) {
	salaireBase := 0
	salaireMajore := 0

	if heuresTravaillees <= 35 {
		salaireBase = heuresTravaillees * tauxHoraire
		salaireMajore = 0
	} else {
		salaireBase = 35 * tauxHoraire
		heuresSup := heuresTravaillees - 35
		salaireMajore = heuresSup * (tauxHoraire * 2)
	}

	return salaireBase, salaireMajore
}

func main() {
	maBase, monSup := calculerSalaire(42, 20)
	total := maBase + monSup
	fmt.Printf("Salaire de base : %d€, Heures sup : %d€, Total : %d€\n", maBase, monSup, total)
}
```

---

## Cas Pratique : Système de Facturation Pay-As-You-Go

### Contexte

Tu es développeur junior dans une boîte de tech. Le Product Manager te demande de coder le moteur de calcul de facturation pour la fin de mois. L'entreprise vend un service d'API avec un modèle **Pay-As-You-Go**.

### ⚠️ Règle d'or : Fintech et Arrondis

**Tous les calculs d'argent doivent être faits en centimes (entiers int)**. Tu convertiras en Euros uniquement lors de l'affichage final.

Exemple : 1050 centimes = 10.50 €

### Règles Métier

| Règle                  | Détail                                  |
| ---------------------- | --------------------------------------- |
| **Abonnement de base** | 1500 centimes (15,00 €)                 |
| **Forfait inclus**     | 10 000 requêtes gratuites               |
| **Surconsommation**    | Chaque requête au-delà coûte 2 centimes |
| **Statut Premium**     | Réduction de 10% sur le montant total   |
| **Validation**         | Les requêtes négatives sont une erreur  |

### Données à tester (Jeu de données)

| Client | Requêtes | Statut      | Notes                              |
| ------ | -------- | ----------- | ---------------------------------- |
| ID 1   | 8 000    | Non-Premium | En dessous du forfait              |
| ID 2   | 15 000   | Premium     | Surconsommation + réduction        |
| ID 3   | -500     | Non-Premium | ⚠️ Erreur                          |
| ID 4   | 12 000   | Non-Premium | Surconsommation                    |
| ID 5   | 20 000   | Premium     | Grosse surconsommation + réduction |

### Sortie attendue

```
--- DÉBUT DU TRAITEMENT DES FACTURES ---
Client ID 1 - Facture : 15.00 € (Premium: Non)
Client ID 2 - Facture : 90.00 € (Premium: Oui)
Client ID 3 - Erreur : Volume de requêtes invalide.
Client ID 4 - Facture : 15.40 € (Premium: Non)
Client ID 5 - Facture : 270.00 € (Premium: Oui)
----------------------------------------
 Chiffre d'affaires total du mois : 390.40 €
```

### Correction

```go
package main

import "fmt"

// calculerFactureClient applique les règles métier pour un client unique.
// Retourne le montant en centimes et un booléen indiquant si le calcul est valide.
func calculerFactureClient(requetes int, isPremium bool) (int, bool) {
	// 1. Règle de Sécurité : Pas de requêtes négatives
	if requetes < 0 {
		return 0, false
	}

	// 2. Tarif de base (15,00 €)
	montant := 1500

	// 3. Calcul de la surconsommation (au-delà de 10 000 requêtes)
	if requetes > 10000 {
		surplusRequetes := requetes - 10000
		montant += surplusRequetes * 2 // 2 centimes par requête
	}

	// 4. Règle Premium : Réduction de 10%
	if isPremium {
		montant = (montant * 90) / 100
	}

	return montant, true
}

func main() {
	// Variable pour accumuler le Chiffre d'Affaires total (en centimes)
	caTotal := 0

	fmt.Println("--- DÉBUT DU TRAITEMENT DES FACTURES ---")

	// --- CLIENT 1 ---
	m1, ok1 := calculerFactureClient(8000, false)
	if ok1 {
		fmt.Printf("Client ID 1 - Facture : %.2f € (Premium: Non)\n", float64(m1)/100)
		caTotal += m1
	} else {
		fmt.Println("Client ID 1 - Erreur : Volume de requêtes invalide.")
	}

	// --- CLIENT 2 ---
	m2, ok2 := calculerFactureClient(15000, true)
	if ok2 {
		fmt.Printf("Client ID 2 - Facture : %.2f € (Premium: Oui)\n", float64(m2)/100)
		caTotal += m2
	} else {
		fmt.Println("Client ID 2 - Erreur : Volume de requêtes invalide.")
	}

	// --- CLIENT 3 ---
	m3, ok3 := calculerFactureClient(-500, false)
	if ok3 {
		fmt.Printf("Client ID 3 - Facture : %.2f € (Premium: Non)\n", float64(m3)/100)
		caTotal += m3
	} else {
		fmt.Println("Client ID 3 - Erreur : Volume de requêtes invalide.")
	}

	// --- CLIENT 4 ---
	m4, ok4 := calculerFactureClient(12000, false)
	if ok4 {
		fmt.Printf("Client ID 4 - Facture : %.2f € (Premium: Non)\n", float64(m4)/100)
		caTotal += m4
	} else {
		fmt.Println("Client ID 4 - Erreur : Volume de requêtes invalide.")
	}

	// --- CLIENT 5 ---
	m5, ok5 := calculerFactureClient(20000, true)
	if ok5 {
		fmt.Printf("Client ID 5 - Facture : %.2f € (Premium: Oui)\n", float64(m5)/100)
		caTotal += m5
	} else {
		fmt.Println("Client ID 5 - Erreur : Volume de requêtes invalide.")
	}

	// --- AFFICHAGE DU BILAN TOTAL ---
	fmt.Println("----------------------------------------")
	fmt.Printf(" Chiffre d'affaires total du mois : %.2f €\n", float64(caTotal)/100)
}
```

---

## Concepts : Slices et Tableaux

### Différences principales

En JavaScript/TypeScript, les Array `[]` s'agrandissent tout seuls. En Go, c'est plus subtil :

- **Tableaux** : Taille fixe, impossible à changer après création
- **Slices** : Taille dynamique, plus flexible

**Dans la vraie vie, on utilise des Slices 99% du temps.**

### Exemples

#### Slice avec valeurs initiales

```go
villes := []string{"Paris", "Marseille", "Lyon"}
```

#### Ajouter des éléments avec `append`

```go
villes := []string{"Paris"}

// Ajouter un élément
villes = append(villes, "Lyon")

// Ajouter plusieurs éléments
villes = append(villes, "Toulouse", "Nice")
```

---

## Étape 5 : Les Maps (Dictionnaires Clés/Valeurs)

### Concept fondamental

Une Map (dictionnaire) associe des clés à des valeurs. La syntaxe se lit de l'intérieur vers l'extérieur :

```
map[Type_De_La_Clé]Type_De_La_Valeur
```

### Déclaration et Initialisation

Créer une map qui associe le nom d'un utilisateur (string) à son âge (int) :

```go
// Création avec des valeurs initiales
ages := map[string]int{
	"Alice": 28,
	"Bob":   35,
}
```

### Ajouter, Modifier et Supprimer

```go
// Ajouter ou Modifier (si la clé existe déjà, ça l'écrase)
ages["Charlie"] = 42

// Supprimer un élément
delete(ages, "Bob")
```

### Le super-réflexe Go : Le "Comma Ok" 💡

En JS, si tu cherches une clé qui n'existe pas, tu reçois `undefined`.
En Go, si la clé n'existe pas, il renvoie la valeur par défaut (pour un int : 0).

**C'est dangereux !** Comment savoir si la personne a vraiment 0 ans ou si elle n'existe pas ?

Go résout ça avec un **retour double** :

```go
// On récupère la valeur ET un booléen (souvent nommé 'ok')
age, ok := ages["Batman"]

if ok {
	fmt.Printf("Batman a %d ans\n", age)
} else {
	fmt.Println("Batman n'est pas dans notre base de données !")
}
```

### Exercice : Le Gestionnaire de Stock

#### Énoncé

Tu vas coder un mini-système de gestion de stock pour un magasin de tech.

**Consignes :**

1. Crée une map nommée `stock` où la clé est le nom du produit (string) et la valeur est la quantité (int)
2. Initialise avec deux produits : "iPhone" (15 unités) et "iPad" (8 unités)
3. Ajoute un nouveau produit : "MacBook" (5 unités)
4. Utilise le système du "Comma ok" pour vérifier si "AppleWatch" est disponible
5. Si oui, affiche sa quantité ; si non, affiche "Désolé, nous ne vendons pas d'AppleWatch."

#### Correction

```go
package main

import "fmt"

func main() {
	// Création de la map avec deux produits
	stock := map[string]int{
		"iPhone": 15,
		"iPad":   8,
	}

	// Ajouter un nouveau produit
	stock["MacBook"] = 5

	// Vérifier si "AppleWatch" existe avec "Comma ok"
	quantite, ok := stock["AppleWatch"]

	if ok {
		fmt.Printf("AppleWatch : %d unités disponibles\n", quantite)
	} else {
		fmt.Println("Désolé, nous ne vendons pas d'AppleWatch.")
	}

	// Afficher tout le stock
	fmt.Println("\n--- Stock Complet ---")
	for produit, qty := range stock {
		fmt.Printf("%s : %d unités\n", produit, qty)
	}
}
```

---

## À Retenir

### Concepts Fondamentaux

✅ **Conditions** — Maîtriser `if`, `else if`, `else`  
✅ **Boucles** — Utiliser `for` avec ses variantes  
✅ **Retours Multiples** — Les tuples (int, bool), (int, int)  
✅ **Validation** — Toujours vérifier les données d'entrée  
✅ **Fintech** — Calculer en centimes pour éviter les arrondis

### Collections

✅ **Slices** — Tableaux dynamiques (utiliser 99% du temps)  
✅ **Append** — Ajouter des éléments aux slices  
✅ **Maps** — Dictionnaires clés/valeurs  
✅ **Comma Ok** — Vérifier l'existence d'une clé avec `value, ok := map[key]`

### Bonnes Pratiques

✅ Nommer les variables de manière explicite  
✅ Ajouter des commentaires sur le code complexe  
✅ Tester avec plusieurs cas (normal, erreur, limite)  
✅ Formater le code avec une indentation cohérente  
✅ Utiliser les structures plutôt que les maps pour du code structuré

---

Jusqu'à présent, nos Slices et nos Maps ne pouvaient stocker qu'un seul type d'information à la fois (uniquement des int, ou uniquement des string). Mais dans la vraie vie d'une entreprise, les données sont hétérogènes. Par exemple, un Utilisateur possède un nom (string), un âge (int), et un statut actif ou non (bool).

C'est là qu'interviennent les Structs. Elles permettent de créer ton propre type de donnée personnalisé, un "moule" qui regroupe plusieurs variables (qu'on appelle des champs).

🟡 Étape 6 : Les Structs (Structures)

1. Déclarer une Struct
   La déclaration se fait toujours en dehors du main (au niveau global du fichier). On utilise les mots-clés type et struct :

Go
type Utilisateur struct {
Nom string
Age int
Actif bool
} 2. Créer une instance (un objet)
Une fois le moule créé, on peut fabriquer autant d'utilisateurs qu'on veut dans notre main :

Go
// Méthode la plus lisible en nommant les champs
u1 := Utilisateur{
Nom: "Alice",
Age: 28,
Actif: true,
} 3. Lire et Modifier les champs
Pour accéder aux informations ou les modifier, on utilise le point . :

Go
// Lire
fmt.Println(u1.Nom) // Affiche: Alice

// Modifier
u1.Age = 29 // Alice fête son anniversaire !
💻 Ton Exercice : La Fiche Produit
Tu vas créer le système de gestion des articles de notre e-commerce.

Consignes :

En dehors de ton main, crée une struct nommée Produit avec trois champs :

Nom (string)

Prix (int en centimes, comme d'habitude !)

Stock (int)

Dans ton main, crée un produit nommé monOrdi avec les valeurs suivantes : "MacBook Pro", un prix de 150000 centimes (1500 €), et un stock initial de 5.

Simule une vente : diminue le stock de monOrdi de 1 unité (en utilisant le point .).

Affiche proprement dans le terminal le nom du produit, son prix formaté en Euros (%.2f) et son nouveau stock.
