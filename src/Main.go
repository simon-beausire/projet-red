package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func concactTabToString(tab []string) string {
	r := ""
	for index, mot := range tab {
		if index != 0 {
			r += " "
		}
		r += string(mot)
	}
	return r
}
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
		if u.nbrMort == 3 {
			fmt.Println("Vous etes mort DEFINITIVEMENT")
			return
		}
		println("vous etes Mort vos points de vie maximum sont diviser par 2")
		u.nbrMort += 1
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
			} else {
				fmt.Println("La potion de poison a bien été utilisé")
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
				if len(split) == 2 {
					i += 1
					fmt.Println("veuillez indiquer le monstre cibler")
				} else {
					index, _ := strconv.Atoi(split[2])
					u.degatsMonstre("Coup de poing", tabSalle, index-1, 2)
				}
			case "bouledefeu":
				if !u.verifierskill("bouledefeu") {
					fmt.Println("Vous ne possedez pas le skill boule de feu")
					i += 1
				} else if len(split) == 2 {
					i += 1
					fmt.Println("veuillez indiquer le monstre cibler")
				} else {
					index, _ := strconv.Atoi(split[2])
					u.degatsMonstre("Coup de poing", tabSalle, index-1, 5)
				}
			}
		case "spellbook":
			if concactTabToString(split[1:]) == "bouledefeu" {
				if u.inInventaire("sort boule de feu") == -1 {
					fmt.Println("vous devez vous procurer le sort auprès du marchand")
					i += 1
				} else {
					u.spellBook(concactTabToString(split[1:]))
					fmt.Println("vous avez appris le sort: Boule de feu")
				}
			} else {
				i += 1
			}
		case "equipement":
			if !u.equipement(concactTabToString(split[1:])) {
				i += 1
			}
		case "pass":
		default:
			i += 1

		}
	}
	u.tour += 1
	u.ActionMonstre(tabSalle)
	u.commande(tabSalle)

}
func main() {
	tabSalle := [][]Salles{
		{Salles{"Jardin des oubliés ", false, true, false, false, []Monstre{initMonstre("loup"), initMonstre("troll"), initMonstre("corbeau")}, false}, Salles{"Couloir nord", false, true, false, true, []Monstre{initMonstre("sanglier"), initMonstre("loup"), initMonstre("corbeau")}, false}, Salles{"Grenier Obscur", false, true, true, false, []Monstre{initMonstre("troll"), initMonstre("sanglier"), initMonstre("corbeau")}, false}},
		{Salles{"Couloir ouest", true, false, false, true, []Monstre{initMonstre("loup"), initMonstre("troll"), initMonstre("sanglier")}, false}, Salles{"Fontaine de depart", true, true, true, false, []Monstre{initMonstre("sanglier")}, true}, Salles{"Couloir est", true, false, false, false, []Monstre{initMonstre("corbeau"), initMonstre("sanglier"), initMonstre("loup")}, false}},
		{Salles{"Salle de Musique", false, true, false, true, []Monstre{initMonstre("troll"), initMonstre("loup"), initMonstre("sanglier")}, false}, Salles{"Couloir sud", true, false, true, true, []Monstre{initMonstre("corbeau"), initMonstre("troll"), initMonstre("loup")}, false}, Salles{"Salle du Banquet", false, true, true, false, []Monstre{initMonstre("sanglier"), initMonstre("corbeau"), initMonstre("troll")}, false}},
		{Salles{"Bureau du Maître", true, false, false, true, []Monstre{initMonstre("loup"), initMonstre("sanglier"), initMonstre("corbeau")}, false}, Salles{" Caves Humides", false, false, true, false, []Monstre{initMonstre("troll"), initMonstre("sanglier"), initMonstre("loup")}, false}, Salles{"Chambre d'Amis", true, false, false, false, []Monstre{initMonstre("corbeau"), initMonstre("troll"), initMonstre("sanglier")}, false}},
	}
	Joueur := initCharacter()
	Joueur.Affichage(tabSalle)
	Joueur.commande(&tabSalle)
}
