package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func concactTab(tab1 []Monstre, tab2 []Monstre) []Monstre {
	for _, monstre := range tab2 {
		tab1 = append(tab1, monstre)
	}
	return tab1
}

func SplitWhiteSpaces(s string) []string {
	r := []string(nil)
	temp := ""
	for _, lettre := range s {
		if lettre == 32 || lettre == 11 || lettre == rune('\n') {
			if temp != "" {
				r = append(r, temp)
			}
			temp = ""
		} else {
			temp += string(lettre)
		}
	}
	r = append(r, temp)
	return r
}

func ToLower(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		if rune(s[i]) > 64 && rune(s[i]) < 91 {
			r += string(rune(s[i]) + 32)
		} else {
			r += string(rune(s[i]))
		}
	}
	return r
}
func ToUpper(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		if !(rune(s[i]) > 96 && rune(s[i]) < 123) {
			r += string(rune(s[i]))
		} else {
			r += string(rune(s[i]) - 32)
		}
	}
	return r
}

func (u *User) isDead() {
	if u.PdvActuel <= 0 {
		println("vous etes Mort vos points de vie maximum sont diviser par 2")
		u.PdvMax = u.PdvMax / 2
		u.PdvActuel = u.PdvMax / 2
	}
}

type Salles struct {
	nom        string
	haut       bool
	bas        bool
	gauche     bool
	droite     bool
	monstre    []Monstre
	personnage bool
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

func (u *User) commande(tabSalle *[][]Salles) {
	for i := u.PointAction; i > 0; i-- {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Action restante: ", rune(i), " | Que voulez vous faire ? : ")
		commande, _ := reader.ReadString('\n')
		commande = strings.TrimSpace(commande)
		split := SplitWhiteSpaces(ToLower(commande))
		switch split[0] {
		case "goto":
			u.deplacementJoueur(tabSalle, split[1])
			u.Affichage(*tabSalle)
		case "takepot":
			if u.inInventaire("potion de soin") >= 0 {
				if u.potionSoin() {
					u.AffichageInventaire()
					println("vous avez utiliser une potion de soin")
				} else {
					i += 1
				}
			} else {
				fmt.Println("Vous ne possedez aucune potion de soin dans votre inventaire")
				i += 1
			}
		case "poisonpot":
			if !u.potionpoison(tabSalle) {
				i += 1
				fmt.Println("Vous ne possedez pas de potion de poison")
			}
		case "accessinventory":
			u.AffichageInventaire()
			i += 1
		case "retour":
			u.Affichage(*tabSalle)
			i += 1
		case "displayinfo":
			i += 1
			if u.Classe == "Nain" {
				u.displayInfoNain()
			} else if u.Classe == "Assassin" {
				u.displayInfoAssassin()
			} else {
				u.displayInfoElfe()
			}

		case "quiter":
			return
		case "marchand":
			u.marchand(tabSalle)
			i += 1
		case "forgeron":
			u.forgeron(tabSalle)
			i += 1
		case "skill":
			switch split[1] {
			case "coupdepoing":
				index, _ := strconv.Atoi(split[2])
				u.degatsMonstre("Coup de poing", tabSalle, index-1, 2)
			}
		case "pass":
		default:
			i += 1
		}
	}
	u.ActionMonstre(tabSalle)
	u.commande(tabSalle)

}
func main() {
	tabSalle := [][]Salles{
		{Salles{"", false, false, false, false, []Monstre{}, false}, Salles{"", false, true, false, false, []Monstre{}, false}, Salles{"", false, false, false, false, []Monstre{}, false}},
		{Salles{"", false, false, false, true, []Monstre{}, false}, Salles{"Fontaine de depart", true, true, true, false, []Monstre{initMonstre("orc")}, true}, Salles{"", false, false, false, false, []Monstre{initMonstre("orc")}, false}},
		{Salles{"", false, false, false, false, []Monstre{}, false}, Salles{"", true, false, false, false, []Monstre{}, false}, Salles{"", false, false, false, false, []Monstre{}, false}},
		{Salles{"", false, false, false, false, []Monstre{}, false}, Salles{"", false, false, false, false, []Monstre{}, false}, Salles{"", false, false, false, false, []Monstre{}, false}},
	}
	Joueur := initCharacter()
	Joueur.Affichage(tabSalle)
	Joueur.commande(&tabSalle)
}
