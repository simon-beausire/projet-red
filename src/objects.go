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

func (u *User) potionpoison(tabSalle *[][]Salles) bool {
	if u.inInventaire("potion de poison") == -1 {
		return false
	} else {
		for index := range (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre {
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].monstre[index].poison += 3
		}
	}
	return true
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
	fmt.Println("bienvenue dans mon magasin, que puis-je pour vous ?")
	fmt.Println("----------ACHETER---------")
	fmt.Println("Potion de Soin : 3 pieces ")
	fmt.Println("potion de Poison : 6 pieces ")
	fmt.Println("Sort Boule de Feu : 25 pieces ")
	fmt.Println("UpgradeInventaire : 30 pieces ")
	fmt.Println("Clé du Donjon : 100 pieces ")
	fmt.Println("----------VENDRE----------")
	fmt.Println("Plume de Corbeau: 1 piece")
	fmt.Println("Cuir de sanglier: 3 piece")
	fmt.Println("Fourrure de Loup: 4 piece")
	fmt.Println("Peau de troll: 7 piece")
	fmt.Println("Pour acheter un objet: acheter nomobjet")
	fmt.Println("Pour vendre un objet: vendre nomobjet")
	fmt.Print("Que voulez vous faire ? : ")
	for true {
		reader := bufio.NewReader(os.Stdin)
		commande, _ := reader.ReadString('\n')
		commande = strings.TrimSpace(commande)
		split := SplitWhiteSpaces(ToLower(commande))
		switch split[0] {
		case "acheter":
			switch split[1] {
			case "upgradeinventaire":
				fmt.Println("aa")
				if u.upgradeInventorySlot() {
					fmt.Println("Vous avez acheter un augmentation de la capacité maximale de l'inventaire pour 30 pieces")
				} else {
					fmt.Println("Vous avez atteint le nombre maximal d'augmentation du nombre de slots")
				}
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

func (u *User) forgeron(tabSalle *[][]Salles) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("bienvenue dans ma forge, que puis-je pour vous ?")
	fmt.Println("Chapeau de l'aventurier")
	fmt.Println("	-1 Plume de Corbeau")
	fmt.Println("	-1 Cuir de Sanglier")
	fmt.Println("Tunique de l'aventurier")
	fmt.Println("	-2 fourrure de Loup")
	fmt.Println("	-1 Peau de Troll")
	fmt.Println("Bottes de l'aventurier")
	fmt.Println("	-1 Fourrure de Loup")
	fmt.Println("	-1 Cuir de Sanglier")
	fmt.Print("Que voulez vous faire ? : ")
	for true {
		reader := bufio.NewReader(os.Stdin)
		commande, _ := reader.ReadString('\n')
		commande = strings.TrimSpace(commande)
		commande = ToLower(commande)
		switch commande {
		case "chapeau de l'aventurier":
			if u.inInventaire("plume de corbeau") != -1 && u.InventaireJoueur[u.inInventaire("plume de corbeau")].Quantite >= 1 && u.inInventaire("cuir de sanglier") != -1 && u.InventaireJoueur[u.inInventaire("cuir de sanglier")].Quantite >= 1 {
				u.retirerInventaire("plume de corbeau", 1)
				u.retirerInventaire("cuir de sanglier", 1)
				if u.ajouterInventaire("chapeau de l'aventurier", 1) {
					fmt.Println("Vous avez fabriquer 1 chapeau de l'aventurier en echange de: 1 cuir de sanglier et 1 plume de corbeau")
				} else {
					fmt.Println("Pas assez de place dans l'inventaire")
				}
			} else {
				fmt.Println("Ressources insuffisante")
			}
		case "tunique de l'aventurier":
			if u.inInventaire("fourrure de loup") != -1 && u.InventaireJoueur[u.inInventaire("fourrure de loup")].Quantite >= 1 && u.inInventaire("peau de troll") != -1 && u.InventaireJoueur[u.inInventaire("peau de troll")].Quantite >= 1 {
				u.retirerInventaire("fourrure de loup", 1)
				u.retirerInventaire("peau de troll", 1)
				if u.ajouterInventaire("tunique de l'aventurier", 1) {
					fmt.Println("Vous avez fabriquer 1 tunique de l'aventurier en echange de: 1 fourrure de loup et 1 peau de troll")
				} else {
					fmt.Println("Pas assez de place dans l'inventaire")
				}
			} else {
				fmt.Println("Ressources insuffisante")
			}
		case "bottes de l'aventurier":
			if u.inInventaire("fourrure de loup") != -1 && u.InventaireJoueur[u.inInventaire("fourrure de loup")].Quantite >= 1 && u.inInventaire("cuir de sanglier") != -1 && u.InventaireJoueur[u.inInventaire("cuir de sanglier")].Quantite >= 1 {
				u.retirerInventaire("fourrure de loup", 1)
				u.retirerInventaire("cuir de sanglier", 1)
				if u.ajouterInventaire("tunique de l'aventurier", 1) {
					fmt.Println("Vous avez fabriquer 1 bottes de l'aventurier en echange de: 1 fourrure de loup et 1 cuir de sanglier")
				} else {
					fmt.Println("Pas assez de place dans l'inventaire")
				}
			} else {
				fmt.Println("Ressources insuffisante")
			}
		case "retour":
			u.Affichage(*tabSalle)
			u.commande(tabSalle)
		}
	}
}
