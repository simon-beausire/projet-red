package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PieceEquipement struct {
	tete       string
	bonusTete  int
	torse      string
	bonusTorse int
	pieds      string
	bonusPied  int
}

type Inventaire struct {
	NomObjet string
	Quantite int
}

type User struct {
	Nom               string
	Classe            string
	Niveau            int
	PdvMax            int
	PdvActuel         int
	InventaireJoueur  []Inventaire
	PlaceInventaire   int
	MaxInventaire     int
	emplacementJoueur []int
	argentJoueur      int
	PointAction       int
	Equipement        PieceEquipement
	tour              int
	skill             []string
	nbrMort           int
}

func (u User) inInventaire(NomObjet string) int {
	for index, objet := range u.InventaireJoueur {
		if objet.NomObjet == NomObjet {
			return index
		}
	}
	return -1
}

func (u *User) ajouterInventaire(NomObjet string, Quantite int) bool {
	if u.PlaceInventaire < u.MaxInventaire {
		if u.inInventaire(NomObjet) == -1 {
			u.InventaireJoueur = append(u.InventaireJoueur, Inventaire{NomObjet, Quantite})
			u.PlaceInventaire += Quantite
		} else {
			u.InventaireJoueur[u.inInventaire(NomObjet)].Quantite += Quantite
			u.PlaceInventaire += Quantite
		}
		return true
	}
	return false
}
func (u *User) retirerInventaire(NomObjet string, Quantite int) {
	if u.inInventaire(NomObjet) == -1 {
		fmt.Println("Action Impossible: l'Objet n'est pas présent dans l'inventaire")
	} else if u.InventaireJoueur[u.inInventaire(NomObjet)].Quantite == 1 {
		u.InventaireJoueur = append(u.InventaireJoueur[:u.inInventaire(NomObjet)], u.InventaireJoueur[(u.inInventaire(NomObjet)+1):]...)
		u.PlaceInventaire -= Quantite
	} else {
		u.InventaireJoueur[u.inInventaire(NomObjet)].Quantite = u.InventaireJoueur[u.inInventaire(NomObjet)].Quantite - 1
		u.PlaceInventaire -= Quantite
	}
}
func (u *User) upgradeInventorySlot() bool {
	if u.MaxInventaire != 40 {
		u.MaxInventaire += 10
		return true
	} else {
		return false
	}

}
func (u *User) deplacementJoueur(tabSalle *[][]Salles, direction string) {
	switch direction {
	case "gauche":
		if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].gauche {
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].personnage = false
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]-1].personnage = true
			u.emplacementJoueur = []int{u.emplacementJoueur[0], u.emplacementJoueur[1] - 1}
		} else {
			fmt.Println("Impossible d'aller a gauche")
			u.commande(tabSalle)
		}
	case "droite":
		if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].droite {
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].personnage = false
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]+1].personnage = true
			u.emplacementJoueur = []int{u.emplacementJoueur[0], u.emplacementJoueur[1] + 1}
		} else {
			fmt.Println("Impossible d'aller a droite")
			u.commande(tabSalle)
		}
	case "haut":
		if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].haut {
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].personnage = false
			(*tabSalle)[u.emplacementJoueur[0]-1][u.emplacementJoueur[1]].personnage = true
			u.emplacementJoueur = []int{u.emplacementJoueur[0] - 1, u.emplacementJoueur[1]}
		} else {
			fmt.Println("Impossible d'aller en haut")
			u.commande(tabSalle)
		}
	case "bas":
		if (*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].bas {
			(*tabSalle)[u.emplacementJoueur[0]][u.emplacementJoueur[1]].personnage = false
			(*tabSalle)[u.emplacementJoueur[0]+1][u.emplacementJoueur[1]].personnage = true
			u.emplacementJoueur = []int{u.emplacementJoueur[0] + 1, u.emplacementJoueur[1]}
		} else {
			fmt.Println("Impossible d'aller en bas")
			u.commande(tabSalle)
		}
	default:
		fmt.Println("Direction non reconnu")
		u.commande(tabSalle)
	}

}

func (u *User) equipement(equipement string) bool {
	if u.inInventaire(equipement) > -1 {
		switch equipement {
		case "chapeau de l'aventurier":
			if u.Equipement.tete == "" {
				u.Equipement.tete = equipement
				u.retirerInventaire("chapeau de l'aventurier", 1)
			} else {
				u.ajouterInventaire(u.Equipement.tete, 1)
				u.PdvMax -= u.Equipement.bonusTete
				u.Equipement.tete = equipement
				u.retirerInventaire("chapeau de l'aventurier", 1)
			}
			fmt.Println(equipement, "Ajouter a votre equipement ")
			u.PdvMax += 10
			u.Equipement.bonusTete = 10
		case "tuniquedel'aventurier":
			if u.Equipement.tete == "" {
				u.Equipement.tete = equipement
				u.retirerInventaire("tunique de l'aventurier", 1)
			} else {
				u.ajouterInventaire(u.Equipement.tete, 1)
				u.PdvMax -= u.Equipement.bonusTorse
				u.Equipement.tete = equipement
				u.retirerInventaire("tunique de l'aventurier", 1)
			}
			fmt.Println(equipement, "Ajouter a votre equipement ")
			u.PdvMax += 25
			u.Equipement.bonusTete = 25
		case "bottesdel'aventurier":
			if u.Equipement.tete == "" {
				u.Equipement.tete = equipement
				u.retirerInventaire("bottes de l'aventurier", 1)
			} else {
				u.ajouterInventaire(u.Equipement.tete, 1)
				u.PdvMax -= u.Equipement.bonusPied
				u.Equipement.tete = equipement
				u.retirerInventaire("bottes de l'aventurier", 1)
			}
			fmt.Println(equipement, "Ajouter a votre equipement ")
			u.PdvMax += 15
			u.Equipement.bonusTete = 15
		default:
			return false
		}
	} else {
		fmt.Println("Vous ne possedez pas cette objet dans votre inventaire")
		return false
	}
	return true
}

func (u *User) ajoutervie(pvajoute int) {
	u.PdvActuel += pvajoute
}

func (u *User) enlevervie(pvenlever int) {
	u.PdvActuel -= pvenlever
	u.isDead()
}

func (u *User) spellBook(sort string) {
	switch sort {
	case "bouledefeu":
		u.skill = append(u.skill, sort)
	}
}

func (u *User) verifierskill(sort string) bool {
	for index := range u.skill {
		if u.skill[index] == sort {
			return true
		}
	}
	return false
}
func initCharacter() User {
	fmt.Print("\033[H\033[2J")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrez votre nom : ")
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)
	if nom == "" {
		return initCharacter()
	}
	nom = ToUpper(string(nom[0])) + ToLower((string(nom[1:])))
	fmt.Print("\033[H\033[2J")
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("classes :")
	fmt.Println("Nain : Resistant au combat")
	fmt.Println("Assassin : Plus agile de son arme, il dispose d'un point d'action suplémentaire")
	fmt.Println("Elf : De part ces connaissances en chimie, c'est le seul a disposer de potion de soin au début du jeu")
	fmt.Print("Choissisez votre classe : ")
	classe, _ := reader.ReadString('\n')
	classe = strings.TrimSpace(classe)
	classe = ToLower(classe)
	switch classe {
	case "nain":
		return User{nom, "Nain", 1, 120, 60, []Inventaire{}, 0, 10, []int{1, 1}, 100, 2, PieceEquipement{}, 0, []string{}, 0}
	case "assassin":
		return User{nom, "Assassin", 1, 100, 50, []Inventaire{}, 0, 10, []int{1, 1}, 100, 4, PieceEquipement{}, 0, []string{}, 0}
	default:
		return User{nom, "elf", 1, 80, 40, []Inventaire{Inventaire{"potion de soin", 3}}, 1, 10, []int{1, 1}, 100, 2, PieceEquipement{}, 0, []string{}, 0}
	}
}
