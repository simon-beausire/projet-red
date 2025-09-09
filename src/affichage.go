package main

import "fmt"

func Espace(esp int, chaine1 string, chaine2 string) string {
	esp -= (len(chaine1) + len(chaine2))
	for i := 0; i <= esp; i++ {
		chaine1 += " "
	}
	return chaine1 + chaine2 + string(rune(127))
}

type Salle struct {
	ligne1  string
	ligne2  string
	ligne3  string
	ligne4  string
	ligne5  string
	ligne6  string
	ligne7  string
	ligne8  string
	ligne9  string
	ligne10 string
}

func (u User) Affichage() {
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
			fmt.Println("|", TabStat[i][index], "|")
		}
	}
}
