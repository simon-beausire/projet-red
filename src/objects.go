package main

import (
	"fmt"
)

func (u *User) potionSoin() {
	valeur := "potion_de_soin"
	for i, v := range u.InventaireJoueur {
		if v.NomObjet == valeur {
			u.ajoutervie(50)
			u.InventaireJoueur = append(u.InventaireJoueur[:i], u.InventaireJoueur[i+1:]...)

		}
	}
}
func (m *Monstre) empoisonerMonstre() {
	m.enlevervie2(3)
}

func (u *User) potionpoison(empoisonerMonstre func()) {
	valeur := "potion_de_poison"
	for i, v := range u.InventaireJoueur {
		if v.NomObjet == valeur {
			empoisonerMonstre()
			u.InventaireJoueur = append(u.InventaireJoueur[:i], u.InventaireJoueur[i+1:]...)
		}
	}
}

func (u *User) forgeron() {
	fmt.Print("\033[H\033[2J")
	println("bienvenue dans mon magasin, que puije pour vous ?")
	fmt.Println("Clé du Donjon : 100 pieces ")

}
