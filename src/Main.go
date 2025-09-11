package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func commande() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Que voulez vous faire ? : ")
	commande, _ := reader.ReadString('\n')
	commande = strings.TrimSpace(commande)
	split := SplitWhiteSpaces(ToLower(commande))
	fmt.Println(split[0])
}
func main() {
	Joueur := initCharacter()
	fmt.Println(Joueur.InventaireJoueur)
	Joueur.ajouterInventaire("Potion de Soin", 2)
	Joueur.Affichage()
	commande()
}
