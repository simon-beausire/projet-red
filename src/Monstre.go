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
	if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].bas {
		(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre = concactTab((*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre, (*tabSalle)[u.emplacementJoueur[0]+1][u.emplacementJoueur[1]].monstre)
		(*tabSalle)[u.emplacementJoueur[0]+1][u.emplacementJoueur[1]].monstre = []Monstre{}
	}
	if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].haut {
		(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre = concactTab((*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre, (*tabSalle)[u.emplacementJoueur[0]-1][u.emplacementJoueur[1]].monstre)
		(*tabSalle)[u.emplacementJoueur[0]-1][u.emplacementJoueur[1]].monstre = []Monstre{}
	}
	if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].gauche {
		(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre = concactTab((*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre, (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]-1].monstre)
		(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]-1].monstre = []Monstre{}
	}
	if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].droite {
		(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre = concactTab((*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre, (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]+1].monstre)
		(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]+1].monstre = []Monstre{}
	}
	u.Affichage(*tabSalle)
	for _, monstre := range (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre {
		u.attaqueMonstre(monstre)
	}
}

func (u User) attaqueMonstre(monstre Monstre) {
	switch monstre.Espece {
	case "orc":
		u.enlevervie(2)
		fmt.Println("Un Orc vous inflige 2 de dégats")
	case "gobelin":
		u.enlevervie(1)
		fmt.Println("Un gobelin vous inflige 1 de dégats")
	case "elfs":
		u.enlevervie(3)
		fmt.Println("Un elf vous inflige 3 de dégats")
	}
}
