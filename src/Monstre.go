package main

import "fmt"

type Monstre struct {
	Espece      string
	PdvMax      int
	PdvActuel   int
	PointAction int
}

func initMonstre(espece string) Monstre {
	switch espece {
	case "orc":
		return Monstre{"orc", 5, 5, 2}
	case "gobelin":
		return Monstre{"gobelin", 9, 9, 3}
	default:
		return Monstre{"elfs", 2, 2, 1}
	}
}

func (u User) ActionMonstre(tabSalle *[][]Salles) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			for _, monstre := range (*tabSalle)[i][j].monstre {
				for k := 0; k < monstre.PointAction; k++ {
					if i == u.emplacementJoueur[0] && j == u.emplacementJoueur[1] {
						fmt.Print("attaquer")
					} else if i-u.emplacementJoueur[0] < 0 {
						fmt.Println("a")
					} else {
						fmt.Print("b")
					}
				}
			}
		}
	}
}
