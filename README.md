# Mes Exercices Go

## Table des matières

1. [Exercice 1 : Le Filtre Élite](#exercice-1--le-filtre-élite)
2. [Exercice 2 : Le Calculateur de Paie](#exercice-2--le-calculateur-de-paie)
3. [Cas Pratique : Système de Facturation Pay-As-You-Go](#cas-pratique--système-de-facturation-pay-as-you-go)
4. [Concepts : Slices et Tableaux](#concepts--slices-et-tableaux)
5. [Étape 5 : Les Maps (Dictionnaires Clés/Valeurs)](#étape-5--les-maps-dictionnaires-clésvaleurs)
6. [Étape 6 : Les Structs (Structures)](#étape-6--les-structs-structures)
7. [Étape 7 : Les Interfaces](#étape-7--les-interfaces)
8. [À Retenir](#à-retenir)

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

## Étape 6 : Les Structs (Structures)

### Pourquoi les utiliser ?

Jusqu'à présent, les slices et les maps nous permettaient de stocker des données de manière pratique, mais elles sont surtout adaptées à des collections simples. Dans la vraie vie, un objet possède souvent plusieurs informations liées entre elles : un nom, un âge, un statut, un prix, etc.

C’est là qu’interviennent les structs. Une struct permet de créer son propre type de données personnalisé, comme un modèle ou un moule.

### Déclaration d’une struct

```go
type Utilisateur struct {
	Nom  string
	Age  int
	Actif bool
}
```

### Créer une instance

```go
u1 := Utilisateur{
	Nom:   "Alice",
	Age:   28,
	Actif: true,
}
```

### Lire et modifier les champs

```go
fmt.Println(u1.Nom)
u1.Age = 29
```

### Exercice : La Fiche Produit

#### Énoncé

Crée une struct nommée `Produit` avec trois champs :

- `Nom string`
- `Prix int`
- `Stock int`

Dans le `main`, crée un produit nommé `monOrdi` avec les valeurs suivantes :

- `"MacBook Pro"`
- `150000` centimes
- `5` unités de stock

Simule une vente en réduisant le stock de `1` unité puis affiche proprement le nom, le prix en euros et le stock restant.

#### Correction

```go
package main

import "fmt"

type Produit struct {
	Nom   string
	Prix  int
	Stock int
}

func main() {
	monOrdi := Produit{
		Nom:   "MacBook Pro",
		Prix:  150000,
		Stock: 5,
	}

	monOrdi.Stock -= 1

	fmt.Printf("Produit : %s\n", monOrdi.Nom)
	fmt.Printf("Prix : %.2f €\n", float64(monOrdi.Prix)/100)
	fmt.Printf("Stock restant : %d\n", monOrdi.Stock)
}
```

---

## Étape 7 : Les Interfaces

### Concept clé

Une interface décrit un comportement attendu. Elle ne contient pas de données concrètes, mais une liste de méthodes que doit implémenter un type.

L’idée principale : un type satisfait une interface s’il possède les méthodes requises.

### Exemple simple

```go
type Affichable interface {
	Afficher() string
}

type Produit struct {
	Nom string
}

func (p Produit) Afficher() string {
	return p.Nom
}

func afficherDetails(a Affichable) {
	fmt.Println(a.Afficher())
}
```

### Pourquoi les interfaces sont utiles ?

Elles permettent de travailler avec différents types de manière uniforme, sans avoir à connaître leur structure interne.

### Exercice guidé

Crée une interface `Affichable` avec une méthode `Afficher() string`. Ensuite, implémente cette interface pour un type `Produit` et un type `Client`, puis teste la fonction `afficherDetails`.

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

### Modélisation et Abstraction

✅ **Structs** — Regrouper plusieurs champs liés dans un même type  
✅ **Interfaces** — Définir un contrat de comportement  
✅ **Méthodes** — Ajouter des comportements à un type  
✅ **Séparation des responsabilités** — Écrire un code plus propre et plus extensible

### Bonnes Pratiques

✅ Nommer les variables de manière explicite  
✅ Ajouter des commentaires sur le code complexe  
✅ Tester avec plusieurs cas (normal, erreur, limite)  
✅ Formater le code avec une indentation cohérente  
✅ Utiliser les structs et les interfaces pour structurer proprement un projet

**Gestion Utilisateur**

package main
import "fmt"

type GestionnaireUtilisateur struct {
BaseDonnees map[string]string
}

func (g *GestionnaireUtilisateur) Inscrire(pseudo string, mdp string) {
valeur, estPresent := g.BaseDonnees[pseudo]
\_= valeur
if len (mdp) <8 {
fmt.Printf("Erreur : Le mot de passe doit contenir au moins 8 caractères.\n")
return
}
if estPresent {
fmt.Println("Déjà pris")
return
}
g.BaseDonnees[pseudo] = mdp
fmt.Printf("Succès : Utililisateur '%s' est inscrit avec succès\n",pseudo)
}
func (g *GestionnaireUtilisateur) Connexion(pseudo string, mdp string) bool {
mdpStocke, estPresent := g.BaseDonnees[pseudo]

    if !estPresent {
        fmt.Printf("Erreur ! Utililisateur inconnue\n")
        return false
    }
    if mdpStocke != mdp {
        fmt.Printf("Mot de passe incorrect\n")
        return false
    }
    fmt.Printf("Bienvenue %s", pseudo)
    return true

}

func main() {
gestionnaire := GestionnaireUtilisateur{
BaseDonnees : make(map[string]string),
}
fmt.Println("--- Début du scénario")
gestionnaire.Inscrire("demetrius", "secret123")
gestionnaire.Inscrire("demetrius", "autremdp")
gestionnaire.Connexion("charlie","mifml")
gestionnaire.Connexion("demetrius", "dlf")
gestionnaire.Connexion("demetrius", "secret123")
}

🧱 Étape 1 : Les Structures
Tu as besoin de deux structures :

Livre : Contient un Titre (string) et EstEmprunte (bool).

Bibliotheque : Contient une map appelée Catalogue. Cette map associe un string (le code unique du livre, ex: "LIV-01") à une structure Livre.

📜 Étape 2 : Les Méthodes (Avec gestion des erreurs)
Attache deux méthodes à ton Bibliotheque :

AjouterLivre(code string, titre string) error :

Clause de garde : Si le code existe déjà dans la map, renvoyer une erreur : errors.New("le code livre existe déjà").

Action : Sinon, ajouter le livre dans la map (avec EstEmprunte à false par défaut) et renvoyer nil.

EmprunterLivre(code string) error :

Clause de garde 1 : Si le code du livre n'existe pas dans la map, renvoyer : errors.New("livre introuvable dans le catalogue").

Clause de garde 2 : Si le livre existe mais que sa variable EstEmprunte est déjà à true, renvoyer : errors.New("ce livre est déjà emprunté par quelqu'un d'autre").

Action : Si tout est bon, modifier le livre dans la map pour le passer à EstEmprunte = true et renvoyer nil.

(⚠️ Attention : Quand tu récupères une structure depuis une map, Go te donne une COPIE. Pour modifier le livre dans la map, n'oublie pas de le réassigner après modification ! Ex: g.Catalogue[code] = monLivreModifie).

🚀 Étape 3 : Le Scénario de test (Ton main)
Dans ton main :

Crée une bibliothèque et initialise sa map.

Ajoute le livre "Le Seigneur des Anneaux" avec le code "LOTR".

Tente d'ajouter à nouveau un livre avec le code "LOTR" (gère et affiche l'erreur obtenue).

Tente d'emprunter le livre "HARRY" (gère et affiche l'erreur : livre introuvable).

Emprunte le livre "LOTR" (doit réussir, affiche un message de succès).

Tente d'emprunter à nouveau "LOTR" (gère et affiche l'erreur : déjà emprunté).

C'est le test ultime pour manipuler les maps et intégrer la gestion des erreurs officielle de Go. À toi de jouer !

\*\*Gestion Bibliothèque

package main

import (
"errors"
"fmt"
)
type Livre struct {
Titre string
EstEmprunte bool
}
type Biblitoheque struct {
Catalogue map[string]Livre
}
func (b \*Biblitoheque) AjouterLivre(code string, titre string) error {
\_, existe := b.Catalogue[code]
if existe {
return errors.New("Le code livre exis déjà")
}
b.Catalogue[code] = Livre {
Titre: titre,
EstEmprunte:false,
}
return nil

}

func (b \*Biblitoheque) EmprunterLivre(code string) error {
livre, existe := b.Catalogue[code]

    if !existe{
        return errors.New("Livre introuvable dans le catalogue.")

    }
    if livre.EstEmprunte {
        return errors.New("Ce livre est déjà emprunté par quelqu'un d'autre.")
    }
    livre.EstEmprunte = true
    b.Catalogue[code] = livre
    return nil

}

func main() {
biblio :=Biblitoheque {
Catalogue: make(map[string]Livre),
}
err := biblio.AjouterLivre("LOTR", "Le Seigneur des anneaux")

    if err != nil {
        fmt.Printf("Erreur : %v\n", err)
    } else {
        fmt.Println("Livre 'LOTR' ajouté avec succès")
    }
    err = biblio.AjouterLivre("LOTR", "Un autre livre")
    if err != nil {
        fmt.Printf("Erreur : %v\n", err)
    }
    err = biblio.EmprunterLivre("HARRY")
    if err != nil {
        fmt.Printf("Erreur : %v\n", err)

    } else {
        fmt.Printf("Livre HARRYY emprunté avec succès.")
    }
    err = biblio.EmprunterLivre("LOTR")
    if err != nil {
         fmt.Printf("Erreur : %v\n", err)

    }else {
        fmt.Printf("Livre LOTR emprunté avec succès.")
    }
     err = biblio.EmprunterLivre("LOTR")
    if err != nil {
         fmt.Printf("Erreur : %v\n", err)

    }else {
        fmt.Printf("Livre LOTR emprunté avec succès.")
    }

}

package main

import (
"errors"
"fmt"
"os"
"strings"
)

// Gestionnaire contient la map des tâches et le chemin du fichier de sauvegarde
type Gestionnaire struct {
Taches map[string]string
FichierSauvegarde string
}

// AjouterTache ajoute une tâche avec validation stricte
func (g \*Gestionnaire) AjouterTache(id string, description string) error {
// Vérification : description vide
if description == "" {
return errors.New("la description ne peut pas être vide")
}

    // Vérification : ID déjà existant
    if _, existe := g.Taches[id]; existe {
    	return errors.New("l'ID de la tâche existe déjà")
    }

    // Ajout dans la map
    g.Taches[id] = description
    return nil

}

// Sauvegarder écrit les tâches dans le fichier ou renvoie une erreur
func (g \*Gestionnaire) Sauvegarder() error {
// Vérification : map vide
if len(g.Taches) == 0 {
return errors.New("aucune tâche à sauvegarder")
}

    // Conversion en format texte
    var sb strings.Builder
    for id, desc := range g.Taches {
    	sb.WriteString(fmt.Sprintf("%s: %s\n", id, desc))
    }
    contenu := sb.String()

    // Écriture dans le fichier
    err := os.WriteFile(g.FichierSauvegarde, []byte(contenu), 0644)
    if err != nil {
    	return err
    }

    return nil

}

func main() {
// Initialisation
gestionnaire := Gestionnaire{
Taches: make(map[string]string),
FichierSauvegarde: "todo.txt",
}

    // 1. Test : Ajout avec description vide
    err := gestionnaire.AjouterTache("T0", "")
    if err != nil {
    	fmt.Printf("Erreur attendue (description vide) : %v\n", err)
    }

    // 2. Ajout de deux tâches valides
    err = gestionnaire.AjouterTache("T1", "Acheter du pain")
    if err != nil {
    	fmt.Printf("Erreur inattendue : %v\n", err)
    }

    err = gestionnaire.AjouterTache("T2", "Coder en Go")
    if err != nil {
    	fmt.Printf("Erreur inattendue : %v\n", err)
    }

    // 3. Sauvegarde
    err = gestionnaire.Sauvegarder()
    if err != nil {
    	fmt.Printf("Erreur de sauvegarde : %v\n", err)
    } else {
    	fmt.Println("Sauvegarde réussie dans todo.txt")
    }

    // Test bonus : Vérifier la duplication d'ID
    err = gestionnaire.AjouterTache("T1", "Duplication")
    if err != nil {
    	fmt.Printf("Erreur attendue (ID dupliqué) : %v\n", err)
    }


}
package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// GestionnaireLogs gère le cycle de vie d'un dossier de logs.
type GestionnaireLogs struct {
	Dossier string
}

// Initialiser vérifie l'existence du dossier et le crée si nécessaire.
func (g *GestionnaireLogs) Initialiser() error {
	// 1. Vérifier si le dossier existe avec os.Stat
	info, err := os.Stat(g.Dossier)

	// Cas 1 : Le dossier n'existe pas -> On le crée
	if os.IsNotExist(err) {
		// 0755 donne les droits rwx pour le propriétaire, rx pour les autres
		return os.Mkdir(g.Dossier, 0755)
	}

	// Cas 2 : Une autre erreur survient (problème de permission, chemin invalide, etc.)
	if err != nil {
		return err
	}

	// Cas 3 : Le chemin existe, mais est-ce bien un dossier ?
	if !info.IsDir() {
		return errors.New("le chemin spécifié existe mais n'est pas un dossier")
	}

	// Le dossier existe déjà et est valide
	return nil
}

// CreerFichierLog crée un fichier de log avec le contenu fourni.
func (g *GestionnaireLogs) CreerFichierLog(nomFichier string, contenu string) error {
	// Validation stricte du contenu
	if contenu == "" {
		return errors.New("contenu vide interdit")
	}

	// Construction du chemin complet de manière sécurisée
	cheminComplet := filepath.Join(g.Dossier, nomFichier)

	// Écriture du fichier
	// 0644 donne les droits rw pour le propriétaire, r pour les autres
	return os.WriteFile(cheminComplet, []byte(contenu), 0644)
}

// PurgerDossier supprime entièrement le dossier de logs et son contenu.
func (g *GestionnaireLogs) PurgerDossier() error {
	// os.RemoveAll supprime le chemin et tout ce qu'il contient (fichiers, sous-dossiers)
	// Si le dossier n'existe déjà pas, os.RemoveAll ne renvoie pas d'erreur (comportement standard Go)
	return os.RemoveAll(g.Dossier)
}

func main() {
	g := GestionnaireLogs{
		Dossier: "mes_logs",
	}

	// 1. Initialisation
	g.Initialiser()

	// 2. Test contenu vide
	err := g.CreerFichierLog("test.log", "")
	if err == nil {
		fmt.Println("Erreur : aurait dû refuser le contenu vide !")
	}

	// 3. Création d'un log valide
	g.CreerFichierLog("app.log", "2026-07-21 : Systeme demarre avec succes")

	// 4. Purge du dossier
	err = g.PurgerDossier()
	if err != nil {
		fmt.Println("Erreur lors de la purge :", err)
	} else {
		fmt.Println("Dossier de logs purgé avec succès !")
	}
}

