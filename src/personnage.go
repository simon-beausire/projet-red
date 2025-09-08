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
}

func initCharacter() User {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Entrez votre nom : ")
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)
	fmt.Println("Bienvenue", nom)
	return User{nom, "humain", 1, 9, 9, []Inventaire{}}
}
