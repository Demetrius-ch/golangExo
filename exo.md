Voici ton sujet d'exercice complet pour la Gestion de Bibliothèque.Cet exercice regroupe tout ce que tu as appris jusqu'à présent : le CRUD complet (GET, POST, PUT, DELETE), la fonction init() pour les IDs dynamiques, l'architecture propre avec switch r.Method, et le middleware de log.
Exercice : L'API "Gestion de Bibliothèque"
📜 Le Cahier des Charges1. La Structure LivreCrée une structure Livre avec les champs suivants (et leurs tags JSON correspondant) :ID (int) $\rightarrow$ "id"Titre (string) $\rightarrow$ "titre"Auteur (string) $\rightarrow$ "auteur"Disponible (bool) $\rightarrow$ "disponible"2. Données initiales & InitialisationDéclare une slice livres pré-remplie avec au moins 2 livres (par exemple ID 1 et ID 2).Déclare une variable globale dernierID int.Implémente la fonction init() qui parcourt la slice livres au démarrage pour trouver l'ID le plus élevé et initialiser dernierID.
3. Les Handlers HTTP (/api/livres)GET (handleGetLivres) :Si aucun paramètre id n'est fourni $\rightarrow$ Renvoie la liste complète des livres au format JSON (200 OK).Si ?id=X est fourni $\rightarrow$ Recherche et renvoie le livre correspondant (200 OK) ou une erreur {"erreur": "Livre non trouvé"} (404 Not Found).POST (handleCreateLivre) :Décode le body JSON reçu.Validation : Le Titre et l'Auteur ne doivent pas être vides. Sinon $\rightarrow$ 400 Bad Request.Règle métier : Lors de la création d'un livre, son champ Disponible doit toujours être mis à true par défaut.Incrémente dernierID, l'attribue au nouveau livre, l'ajoute à la slice, et renvoie le livre créé (201 Created).PUT (handleUpdateLivre) :Vérifie que le paramètre ?id=X est présent dans l'URL.Décode le body JSON (nouveau titre, nouvel auteur, statut disponible).Met à jour le livre dans la slice livres.Renvoie le livre mis à jour (200 OK) ou une erreur 404 Not Found si l'ID n'existe pas.DELETE (handleDeleteLivre) :Vérifie que le paramètre ?id=X est présent dans l'URL.Supprime le livre de la slice grâce au slicing (append(livres[:i], livres[i+1:]...)).Renvoie le livre supprimé (200 OK) ou 404 Not Found.4. Le Handler Principal & MiddlewareCrée la fonction handlerLivres(w, r) avec le switch r.Method.Ajoute le middleware loggerMiddleware qui affiche [METHODE] URL dans le terminal.Lance le serveur sur le port :8080.🧪 Commandes curl pour tester ton code une fois écrit :Bash# 1. Obtenir tous les livres
curl -X GET http://localhost:8080/api/livres

# 2. Obtenir le livre ID 1

curl -X GET "http://localhost:8080/api/livres?id=1"

# 3. Ajouter un nouveau livre

curl -X POST http://localhost:8080/api/livres \
 -H "Content-Type: application/json" \
 -d '{"titre": "L'Étranger", "auteur": "Albert Camus"}'

# 4. Modifier le statut d'un livre (ex: le rendre indisponible)

curl -X PUT "http://localhost:8080/api/livres?id=1" \
 -H "Content-Type: application/json" \
 -d '{"titre": "Pot de peinture v2", "auteur": "Auteur X", "disponible": false}'

# 5. Supprimer un livre

curl -X DELETE "http://localhost:8080/api/livres?id=2"
Prends ton temps, écris le code complet dans ton fichier main.go, et colle-le ici dès que tu es prêt. Nous le corrigerons ensemble avant d'attaquer la sauvegarde sur disque JSON !

Exercice 1 : Recherche & Filtrage Dynamique (GET)
Objectif : Enrichir le handler GET /api/livres pour supporter le filtrage par statut de disponibilité.

Consigne :

Si l'URL contient ?disponible=true, renvoie uniquement la liste des livres disponibles.

Si l'URL contient ?disponible=false, renvoie uniquement les livres empruntés.

Si aucun paramètre n'est fourni, renvoie tous les livres.

Exercice 2 : Recherche par Mot-Clé (GET)
Objectif : Implémenter une recherche textuelle simple sur l'auteur ou le titre.

Consigne :

Accepte un paramètre ?q=mot_cle (ex: /api/livres?q=camus).

La recherche doit être insensible à la casse (ignore les majuscules/minuscules via strings.ToLower).

Renvoie la liste des livres dont le titre ou l'auteur contient le mot-clé.

Exercice 3 : Emprunter / Rendre un Livre (PATCH ou POST)
Objectif : Ajouter des routes d'action métier spécifiques plutôt qu'un PUT générique.

Consigne :

Crée la route POST /api/livres/emprunter?id=X. Si le livre est disponible, passe son statut disponible à false. S'il est déjà emprunté, renvoie une erreur 400 Bad Request avec le message "Livre déjà emprunté".

Crée la route POST /api/livres/rendre?id=X qui remet le statut à true.

Exercice 4 : Validation Stricte des Données (POST / PUT)
Objectif : Renforcer la sécurité du découpage des données JSON.

Consigne :

Lors de la création ou modification, vérifie que :

Le titre contient au moins 3 caractères.

L'auteur contient au moins 2 caractères.

Si l'une des règles échoue, renvoie un statut 422 Unprocessable Entity avec un message d'erreur explicite.

Exercice 5 : Middleware de Statistiques & Performance
Objectif : Mesurer le temps d'exécution de tes handlers API.

Consigne :

Crée un middleware metricsMiddleware.

À chaque requête, calcule le temps écoulé en millisecondes (time.Since(début)).

Affiche dans la console : [GET] /api/livres - Traité en 0.45ms.

🚀 Le Projet : "Library-API Microservice"
Objectif : Assembler un serveur API complet et propre intégrant toute la logique de la bibliothèque.

📜 Cahier des Charges du Projet
Modèle de données :

Go
type Livre struct {
ID int `json:"id"`
Titre string `json:"titre"`
Auteur string `json:"auteur"`
Disponible bool `json:"disponible"`
}
Routes HTTP à implémenter :

GET /api/livres (supporte ?id=X, ?disponible=true/false, et ?q=recherche)

POST /api/livres (création avec Disponible = true par défaut)

PUT /api/livres?id=X (modification complète)

DELETE /api/livres?id=X (suppression)

POST /api/livres/emprunter?id=X (action métier d'emprunt)

Exigences Techniques :

Initialisation dynamique de dernierID via init().

Complexité cognitive basse (fonctions bien découpées).

Middleware de Logging + Timing sur toutes les routes.

Réponses JSON systématiques avec les bons codes d'état HTTP (200, 201, 400, 404, 422).

Prends ton temps pour coder le projet ou l'exercice de ton choix dans ton IDE, et partage ton code ici quand tu souhaites une revue globale !
