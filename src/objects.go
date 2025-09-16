package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (u *User) potionSoin() bool {
	if u.PdvActuel != u.PdvMax {
		u.ajoutervie(50)
		u.retirerInventaire("potion de soin", 1)
		if u.PdvActuel > u.PdvMax {
			u.PdvActuel = u.PdvMax
		}
		return true
	}
	fmt.Println("Vie deja au maximum")
	return false
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

func (u *User) testAchat(prix int, NomObjet string) {
	if u.argentJoueur < prix {
		fmt.Println("Solde Insuffisant")
	} else if !u.ajouterInventaire(NomObjet, 1) {
		fmt.Println("Plus de place dans l'inventaire")
	} else {
		u.argentJoueur -= prix
		fmt.Println("Vous avez acheter ", NomObjet, "en echange de ", prix, "pieces")
	}

}

func (u *User) marchand(tabSalle *[][]Salles) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("bienvenue dans mon magasin, que puije pour vous ?")
	fmt.Println("----------ACHETER---------")
	fmt.Println("Potion de Soin : 3 pieces ")
	fmt.Println("potion de Poison : 6 pieces ")
	fmt.Println("Sort Boule de Feu : 25 pieces ")
	fmt.Println("Clé du Donjon : 100 pieces ")
	fmt.Println("----------VENDRE----------")
	fmt.Println("Plume de Corbeau: 1 piece")
	fmt.Println("Cuir de sanglier: 3 piece")
	fmt.Println("Fourrure de Loup: 4 piece")
	fmt.Println("Peau de troll: 7 piece")
	fmt.Println("Pour acheter un objet: acheter nomobjet")
	fmt.Println("Pour vendre un objet: vendre nomobjet")
	for true {
		reader := bufio.NewReader(os.Stdin)
		commande, _ := reader.ReadString('\n')
		commande = strings.TrimSpace(commande)
		split := SplitWhiteSpaces(ToLower(commande))
		switch split[0] {
		case "acheter":
			switch split[1] {
			case "potion":
				if split[2] == "soin" {
					u.testAchat(3, "potion de soin")
				} else if split[2] == "poison" && u.argentJoueur >= 6 {
					u.testAchat(6, "potion de poison")
				}
			case "sort":
				u.testAchat(25, "sort boule de feu")
			case "clé":
				u.testAchat(100, "cle du donjon")
			}
		case "vendre":
			switch split[1] {
			case "plume":
				if u.inInventaire("plume de corbeau") >= 0 {
					u.retirerInventaire("plume de corbeau", 1)
					u.argentJoueur += 1
				}
			case "cuir":
				if u.inInventaire("cuir de sanglier") >= 0 {
					u.retirerInventaire("cuir de sanglier", 1)
					u.argentJoueur += 3
				}
			case "fourrure":
				if u.inInventaire("fourrure de Loup") >= 0 {
					u.retirerInventaire("fourrure de Loup", 1)
					u.argentJoueur += 4
				}
			case "peau":
				if u.inInventaire("Peau de troll") >= 0 {
					u.retirerInventaire("Peau de troll", 1)
					u.argentJoueur += 7
				}
			}
		case "retour":
			u.Affichage(*tabSalle)
			u.commande(tabSalle)
		}
	}
}
