package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SplitWhiteSpaces(s string) []string {
	r := []string(nil)
	temp := ""
	for _, lettre := range s {
		if lettre == 32 || lettre == 11 || lettre == rune('\n') {
			if temp != "" {
				r = append(r, temp)
			}
			temp = ""
		} else {
			temp += string(lettre)
		}
	}
	r = append(r, temp)
	return r
}

func ToLower(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		if rune(s[i]) > 64 && rune(s[i]) < 91 {
			r += string(rune(s[i]) + 32)
		} else {
			r += string(rune(s[i]))
		}
	}
	return r
}
func ToUpper(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		if !(rune(s[i]) > 96 && rune(s[i]) < 123) {
			r += string(rune(s[i]))
		} else {
			r += string(rune(s[i]) - 32)
		}
	}
	return r
}

type Salles struct {
	nom        string
	haut       bool
	bas        bool
	gauche     bool
	droite     bool
	monstre    string
	personnage bool
}

func (u *User) deplacementJoueur(tabSalle *[][]Salles, direction string) {
	switch direction {
	case "gauche":
		if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].gauche {
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].personnage = false
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]-1].personnage = true
			u.emplacementJoueur = []int{u.emplacementJoueur[0], u.emplacementJoueur[1] - 1}
		} else {
			fmt.Println("Impossible d'aller a gauche")
			u.commande(tabSalle)
		}
	case "droite":
		if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].droite {
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].personnage = false
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]+1].personnage = true
			u.emplacementJoueur = []int{u.emplacementJoueur[0], u.emplacementJoueur[1] + 1}
		} else {
			fmt.Println("Impossible d'aller a droite")
			u.commande(tabSalle)
		}
	case "haut":
		if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].haut {
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].personnage = false
			(*tabSalle)[u.emplacementJoueur[0]-1][u.emplacementJoueur[1]].personnage = true
			u.emplacementJoueur = []int{u.emplacementJoueur[0] - 1, u.emplacementJoueur[1]}
		} else {
			fmt.Println("Impossible d'aller en haut")
			u.commande(tabSalle)
		}
	case "bas":
		if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].bas {
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].personnage = false
			(*tabSalle)[u.emplacementJoueur[0]+1][u.emplacementJoueur[1]].personnage = true
			u.emplacementJoueur = []int{u.emplacementJoueur[0] + 1, u.emplacementJoueur[1]}
		} else {
			fmt.Println("Impossible d'aller en bas")
			u.commande(tabSalle)
		}
	default:
		fmt.Println("Direction non reconnu")
		u.commande(tabSalle)
	}

}

func (u *User) commande(tabSalle *[][]Salles) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Que voulez vous faire ? : ")
	commande, _ := reader.ReadString('\n')
	commande = strings.TrimSpace(commande)
	split := SplitWhiteSpaces(ToLower(commande))
	fmt.Println(split[0])
	switch split[0] {
	case "goto":
		u.deplacementJoueur(tabSalle, split[1])
	}
	u.Affichage(*tabSalle)
	u.commande(tabSalle)

}
func main() {
	tabSalle := [][]Salles{
		{Salles{"", false, false, false, false, "", false}, Salles{"", false, true, false, false, "", false}, Salles{"", false, false, false, false, "", false}},
		{Salles{"", false, false, false, true, "", false}, Salles{"Fontaine de depart", true, true, true, false, "", true}, Salles{"", false, false, false, false, "", false}},
		{Salles{"", false, false, false, false, "", false}, Salles{"", true, false, false, false, "", false}, Salles{"", false, false, false, false, "", false}},
		{Salles{"", false, false, false, false, "", false}, Salles{"", false, false, false, false, "", false}, Salles{"", false, false, false, false, "", false}},
	}
	Joueur := initCharacter()
	fmt.Println(Joueur.InventaireJoueur)
	Joueur.ajouterInventaire("Potion de Soin", 2)
	Joueur.Affichage(tabSalle)
	Joueur.commande(&tabSalle)
}
