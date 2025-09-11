package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Inventaire struct {
	NomObjet string
	Quantite int
}

type User struct {
	Nom              string
	Classe           string
	Niveau           int
	PdvMax           int
	PdvActuel        int
	InventaireJoueur []Inventaire
	PlaceInventaire  int
	MaxInventaire    int
}

func (u User) ajouterInventaire(NomObjet string, Quantite int) {
	if u.PlaceInventaire < u.MaxInventaire {
		u.InventaireJoueur = append(u.InventaireJoueur, Inventaire{NomObjet, Quantite})
	}
	fmt.Println(u.InventaireJoueur)
}

func (u *User) ajoutervie(pvajoute int) {
	u.PdvActuel += pvajoute
}

func initCharacter() User {
	fmt.Print("\033[H\033[2J")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrez votre nom : ")
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)
	return User{nom, "humain", 1, 9, 9, []Inventaire{}, 0, 3}
}
