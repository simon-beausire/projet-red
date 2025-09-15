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
	Nom               string
	Classe            string
	Niveau            int
	PdvMax            int
	PdvActuel         int
	InventaireJoueur  []Inventaire
	PlaceInventaire   int
	MaxInventaire     int
	emplacementJoueur []int
	argentJoueur      int
	PointAction       int
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

func (u *User) enlevervie(pvenlever int) {
	u.PdvActuel -= pvenlever
	u.isDead()
}

func (m *Monstre) enlevervie2(pvenlever int) {
	m.PdvActuel -= pvenlever
}

func initCharacter() User {
	fmt.Print("\033[H\033[2J")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrez votre nom : ")
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)
	nom = ToUpper(string(nom[0])) + ToLower((string(nom[1:])))
	fmt.Print("\033[H\033[2J")
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("classes :")
	fmt.Println("Nain : Resistant au combat")
	fmt.Println("Assassin : Plus agile de son arme, il dispose d'un point d'action suplémentaire")
	fmt.Println("Elf : De part ces connaissances en chimie, c'est le seul a disposer de potion de soin au début du jeu")
	fmt.Print("Choissisez votre classe : ")
	classe, _ := reader.ReadString('\n')
	classe = strings.TrimSpace(classe)
	switch classe {
	case "nain":
		return User{nom, "Nain", 1, 15, 15, []Inventaire{}, 0, 3, []int{1, 1}, 10, 2}
	case "assassin":
		return User{nom, "Assassin", 1, 10, 10, []Inventaire{}, 0, 3, []int{1, 1}, 10, 4}
	default:
		return User{nom, "elf", 1, 10, 10, []Inventaire{Inventaire{"potion de soin", 3}}, 0, 3, []int{1, 1}, 10, 2}
	}
}
