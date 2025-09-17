package main

import (
	"fmt"
	"time"
)

type Monstre struct {
	Espece      string
	PdvMax      int
	PdvActuel   int
	PointAction int
	poison      int
}

func initMonstre(espece string) Monstre {
	switch espece {
	case "orc":
		return Monstre{"orc", 5, 5, 2, 0}
	case "gobelin":
		return Monstre{"gobelin", 9, 9, 3, 0}
	default:
		return Monstre{"elfs", 2, 2, 1, 0}
	}
}

func (u *User) ActionMonstre(tabSalle *[][]Salles) {
	if len((*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre) != 3 {
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
		fmt.Println("Des monstres ce déplace sur vous")
		time.Sleep(1 * time.Second)
	}
	for _, monstre := range (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre {
		u.attaqueMonstre(monstre, tabSalle)
	}
}

func (u *User) attaqueMonstre(monstre Monstre, tabSalle *[][]Salles) {
	switch monstre.Espece {
	case "orc":
		u.enlevervie(2)
		u.Affichage(*tabSalle)
		fmt.Println("Un Orc vous inflige 2 de dégats")
	case "gobelin":
		u.enlevervie(1)
		u.Affichage(*tabSalle)
		fmt.Println("Un gobelin vous inflige 1 de dégats")
	case "elfs":
		u.enlevervie(3)
		u.Affichage(*tabSalle)
		fmt.Println("Un elf vous inflige 3 de dégats")
	}
}

func (u *User) coupDePoing(tabSalle *[][]Salles, monstre int) {
	(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre[monstre].PdvActuel -= 2
	fmt.Println("vous infligez 2 points de degats a ", (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre[monstre].Espece)
}
