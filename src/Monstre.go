package main

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
	if u.emplacementJoueur[0] > 0 && (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].haut {
		(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre = concactTab((*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre, (*tabSalle)[u.emplacementJoueur[0]-1][u.emplacementJoueur[1]].monstre)
		(*tabSalle)[u.emplacementJoueur[0]-1][u.emplacementJoueur[1]].monstre = []Monstre{}
		u.Affichage(*tabSalle)
	} else if u.emplacementJoueur[0] > 0 && (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].haut {
		(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre = concactTab((*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre, (*tabSalle)[u.emplacementJoueur[0]-1][u.emplacementJoueur[1]].monstre)
		(*tabSalle)[u.emplacementJoueur[0]-1][u.emplacementJoueur[1]].monstre = []Monstre{}
		u.Affichage(*tabSalle)
	}
}
