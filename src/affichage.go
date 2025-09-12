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

func (u User) Affichage(tabSalle [][]Salles) {

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
			if i%9 == 0 && tabSalle[i/9][j].haut {
				fmt.Print("______________|  |______________")
			} else if i%9 == 0 && !tabSalle[i/9][j].haut {
				fmt.Print("________________________________")
			} else if i%9 == 1 {
				fmt.Print("|", centrer(28, tabSalle[i/9][j].nom), "|")
			} else if i%9 == 2 && tabSalle[i/9][j].personnage {
				fmt.Print("| ¤-->  Vous etes Ici !! <--¤  |")
			} else if i%9 == 2 || i%9 == 3 {
				fmt.Print("|                              |")
			} else if i%9 == 4 && tabSalle[i/9][j].droite && tabSalle[i/9][j].gauche {
				fmt.Print("=          Monstres:           =")
			} else if i%9 == 4 && !tabSalle[i/9][j].droite && tabSalle[i/9][j].gauche {
				fmt.Print("=          Monstres:           |")
			} else if i%9 == 4 && tabSalle[i/9][j].droite && !tabSalle[i/9][j].gauche {
				fmt.Print("|          Monstres:           =")
			} else if i%9 == 4 && !tabSalle[i/9][j].droite && !tabSalle[i/9][j].gauche {
				fmt.Print("|          Monstres:           |")
			} else if i%9 == 5 || i%9 == 6 || i%9 == 7 {
				fmt.Print("|", centrer(28, "temp off"), "|")
			} else if i%9 == 8 && tabSalle[i/9][j].bas {
				fmt.Print("|_____________    _____________|")
			} else if i%9 == 8 && !tabSalle[i/9][j].bas {
				fmt.Print("|______________________________|")
			}
		}
		println("")

	}
}
