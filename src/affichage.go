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

func (u User) Affichage() {

	tab := [][]int{
		{0, 1, 0},
		{0, 0, 0},
		{1, 0, 0},
		{0, 0, 1},
	}

	salles := [][][]string{{
		{"______________________________"},
		{"|    FONTAINE DES ARRIVANTS  |"},
		{"|                            |"},
		{"-                            -"},
		{"                              "},
		{"-                            -"},
		{"|                            |"},
		{"|                            |"},
		{"|____________________________|"},
	}, {
		{"______________________________"},
		{"|     CHEMIN DE TRAVERSE     |"},
		{"|                            |"},
		{"-                            -"},
		{"                              "},
		{"-                            -"},
		{"|                            |"},
		{"|                            |"},
		{"|____________________________|"},
	}}

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

			for _, lettre := range salles[tab[i/9][j]][i%9] {
				fmt.Print(string(lettre))
			}
		}
		println("")
	}
}
