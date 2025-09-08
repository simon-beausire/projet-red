package main

type Inventaire struct {
	NomObjet string
	Quantite int
}

type User struct {
	Nom              string
	Classe           string
	Niveau           string
	PdvMax           int
	PdvActuel        int
	InventaireJoueur []Inventaire
}
