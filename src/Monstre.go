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
	loot        string
}

func (m *Monstre) MonstreisDead(u *User, tabSalle *[][]Salles) bool {
	if m.PdvActuel <= 0 {
		u.Affichage(*tabSalle)
		println(m.Espece, " a succomber a ses blessures")
		return true
	} else {
		return false
	}
}

func initMonstre(espece string) Monstre {
	switch espece {
	case "corbeau":
		return Monstre{"corbeau", 2, 2, 2, 0, "plume de corbeau"}
	case "loup":
		return Monstre{"loup", 3, 3, 3, 0, "fourrure de loup"}
	case "troll":
		return Monstre{"troll", 9, 9, 1, 0, "peau de troll"}
	default:
		return Monstre{"sanglier", 4, 4, 1, 0, "cuir de sanglier"}
	}
}

func (u *User) ActionMonstre(tabSalle *[][]Salles) {
	if len((*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre) != 3 {
		if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].bas {
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre = concactTab((*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre, (*tabSalle)[u.emplacementJoueur[0]+1][u.emplacementJoueur[1]].monstre)
			(*tabSalle)[u.emplacementJoueur[0]+1][u.emplacementJoueur[1]].monstre = []Monstre{}
			fmt.Println("Des monstres ce déplace sur vous")
		}
		if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].haut {
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre = concactTab((*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre, (*tabSalle)[u.emplacementJoueur[0]-1][u.emplacementJoueur[1]].monstre)
			(*tabSalle)[u.emplacementJoueur[0]-1][u.emplacementJoueur[1]].monstre = []Monstre{}
			fmt.Println("Des monstres ce déplace sur vous")
		}
		if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].gauche {
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre = concactTab((*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre, (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]-1].monstre)
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]-1].monstre = []Monstre{}
			fmt.Println("Des monstres ce déplace sur vous")
		}
		if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].droite {
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre = concactTab((*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre, (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]+1].monstre)
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]+1].monstre = []Monstre{}
			fmt.Println("Des monstres ce déplace sur vous")
		}
		u.Affichage(*tabSalle)
		time.Sleep(1 * time.Second)
	}
	for index, monstre := range (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre {
		if monstre.poison > 0 {
			monstre.poison -= 1
			if !u.degatsMonstre("Le poison ", tabSalle, index, 3) {
				break
			}
		}
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

func (u *User) degatsMonstre(attaque string, tabSalle *[][]Salles, index int, degats int) bool {
	(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre[index].PdvActuel -= degats
	if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre[index].MonstreisDead(u, tabSalle) {
		if u.ajouterInventaire((*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre[index].loot, 1) {
			fmt.Println("1 ", (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre[index].loot, "ajouter a votre inventaire")
			u.Niveau += 1
		}
		u.argentJoueur += 15
		(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre = append((*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre[:index], (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre[index+1:]...)
		return false
	}
	u.Affichage(*tabSalle)
	fmt.Println(attaque, " inflige ", degats, " points de degats a ", (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre[index].Espece)
	return true
}
