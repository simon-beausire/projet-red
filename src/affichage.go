package main

import (
	"fmt"
)

func Espace(esp int, chaine1 string, chaine2 string) string {
	esp -= (len(chaine1) + len(chaine2))
	for i := 0; i <= esp; i++ {
		chaine1 += " "
	}
	return chaine1 + chaine2 + string(rune(127))
}
func centrer(esp int, chaine string) string {
	esp = esp - len(chaine)
	r := ""
	for i := 0; i <= esp/2; i++ {
		r += " "
	}
	r += chaine
	for i := 0; i <= esp/2+esp%2; i++ {
		r += " "
	}
	return r
}

func (u User) Affichage() {

	type Salles struct {
		nom        string
		haut       bool
		bas        bool
		gauche     bool
		droite     bool
		monstre    string
		personnage bool
	}

	tabSalle := [][]Salles{
		{Salles{"", false, false, false, false, "", false}, Salles{"", false, false, false, false, "", false}, Salles{"", false, false, false, false, "", false}},
		{Salles{"", false, false, false, false, "", false}, Salles{"Fontaine de depart", false, false, false, false, "", true}, Salles{"", false, false, false, false, "", false}},
		{Salles{"", false, false, false, false, "", false}, Salles{"", false, false, false, false, "", false}, Salles{"", false, false, false, false, "", false}},
		{Salles{"", false, false, false, false, "", false}, Salles{"", false, false, false, false, "", false}, Salles{"", false, false, false, false, "", false}},
	}

	fmt.Print("\033[H\033[2J")
	TabStat := [][]string{
		{"------------------------------"},
		{Espace(29, "Nom Joueur", u.Nom)},
		{Espace(29, "Classe", u.Classe)},
		{Espace(29, "Niveau", string(u.Niveau+48))},
		{Espace(29, "Points de vie ", string(u.PdvActuel+48))},
		{"------------------------------"},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
	}
	for i := 0; i < 36; i++ {
		for index := range TabStat[i] {
			fmt.Print("|", TabStat[i][index], "|")
		}
		for j := 0; j < 3; j++ {
			if i%9 == 0 && tabSalle[i/9][j].haut == true {
				fmt.Print("______________|  |______________")
			} else if i%9 == 0 && tabSalle[i/9][j].haut == false {
				fmt.Print("________________________________")
			} else if i%9 == 1 {
				fmt.Print("|", centrer(28, tabSalle[i/9][j].nom), "|")
			} else if i%9 == 2 && tabSalle[i/9][j].personnage == true {
				fmt.Print("| ¤-->  Vous etes Ici !! <--¤  |")
			} else if i%9 == 2 || i%9 == 3 {
				fmt.Print("|                              |")
			} else if i%9 == 4 && tabSalle[i/9][j].droite == true && tabSalle[i/9][j].gauche == true {
				fmt.Print("=          Monstres:           =")
			} else if i%9 == 4 && tabSalle[i/9][j].droite == false && tabSalle[i/9][j].gauche == true {
				fmt.Print("=          Monstres:           |")
			} else if i%9 == 4 && tabSalle[i/9][j].droite == true && tabSalle[i/9][j].gauche == false {
				fmt.Print("|          Monstres:           =")
			} else if i%9 == 4 && tabSalle[i/9][j].droite == false && tabSalle[i/9][j].gauche == false {
				fmt.Print("|          Monstres:           |")
			} else if i%9 == 5 || i%9 == 6 || i%9 == 7 {
				fmt.Print("|", centrer(28, tabSalle[i/9][j].monstre), "|")
			} else if i%9 == 8 && tabSalle[i/9][j].bas == true {
				fmt.Print("|_____________    _____________|")
			} else if i%9 == 8 && tabSalle[i/9][j].bas == false {
				fmt.Print("|______________________________|")
			}
		}
		println("")

	}
}
