package models

type Livre struct {
	ID         int    `json:"id"`
	Titre      string `json:"titre"`
	Auteur     string `json:"auteur"`
	Disponible bool   `json:"disponible"`
}
