package main

import (
	"fmt"
)

func Espace(esp int, chaine1 string, chaine2 string) string {
	esp -= (len(chaine1) + len(chaine2))
	for i := 0; i <= esp; i++ {
		chaine1 += " "
	}
	return chaine1 + chaine2 + string(rune(127))
}
func centrer(esp int, chaine string) string {
	esp = esp - len(chaine)
	r := ""
	for i := 0; i <= esp/2; i++ {
		r += " "
	}
	r += chaine
	for i := 0; i <= esp/2+esp%2; i++ {
		r += " "
	}
	return r
}

func (u User) Affichage(tabSalle [][]Salles) {
	fmt.Print("\033[H\033[2J")
	niveau := fmt.Sprintf("%d", u.Niveau)
	pdv := fmt.Sprintf("%d", u.PdvActuel)
	tour := fmt.Sprintf("%d", u.tour)
	TabStat := [][]string{
		{"------------------------------"},
		{Espace(29, "Nom Joueur", u.Nom)},
		{Espace(29, "Classe", u.Classe)},
		{Espace(29, "Niveau", niveau)},
		{Espace(29, "Points de vie ", pdv)},
		{Espace(29, "Nombre de tour", tour)},
		{"------------------------------"},
		{"                              "},
		{"                              "},
		{"goto direction                "},
		{"accessinventory               "},
		{"retour                        "},
		{"displayinfo                   "},
		{"quiter                        "},
		{"takepot                       "},
		{"poisonpot                     "},
		{"retour                        "},
		{"marchand                      "},
		{"forgeron                      "},
		{"skill coupdepoing lignemonstre"},
		{"skill bouledefeu lignemonstre "},
		{"pass                          "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
	}
	for i := 0; i < 36; i++ {
		for index := range TabStat[i] {
			fmt.Print("|", TabStat[i][index], "|")
		}
		for j := 0; j < 3; j++ {
			if i%9 == 0 && tabSalle[i/9][j].haut {
				fmt.Print("______________|  |______________")
			} else if i%9 == 0 && !tabSalle[i/9][j].haut {
				fmt.Print("________________________________")
			} else if i%9 == 1 {
				fmt.Print("|", centrer(28, tabSalle[i/9][j].nom), "|")
			} else if i%9 == 2 && tabSalle[i/9][j].personnage {
				fmt.Print("| ¤-->  Vous etes Ici !! <--¤  |")
			} else if i%9 == 2 || i%9 == 3 {
				fmt.Print("|                              |")
			} else if i%9 == 4 && tabSalle[i/9][j].droite && tabSalle[i/9][j].gauche {
				fmt.Print("=          Monstres:           =")
			} else if i%9 == 4 && !tabSalle[i/9][j].droite && tabSalle[i/9][j].gauche {
				fmt.Print("=          Monstres:           |")
			} else if i%9 == 4 && tabSalle[i/9][j].droite && !tabSalle[i/9][j].gauche {
				fmt.Print("|          Monstres:           =")
			} else if i%9 == 4 && !tabSalle[i/9][j].droite && !tabSalle[i/9][j].gauche {
				fmt.Print("|          Monstres:           |")
			} else if i%9 == 5 && len(tabSalle[i/9][j].monstre) > 0 {
				pdv := fmt.Sprintf("%d", tabSalle[i/9][j].monstre[0].PdvActuel)
				fmt.Print("|  1: ", Espace(20, tabSalle[i/9][j].monstre[0].Espece, pdv), "pv  |")
			} else if i%9 == 6 && len(tabSalle[i/9][j].monstre) > 1 {
				pdv := fmt.Sprintf("%d", tabSalle[i/9][j].monstre[1].PdvActuel)
				fmt.Print("|  2: ", Espace(20, tabSalle[i/9][j].monstre[1].Espece, pdv), "pv  |")
			} else if i%9 == 7 && len(tabSalle[i/9][j].monstre) > 2 {
				pdv := fmt.Sprintf("%d", tabSalle[i/9][j].monstre[2].PdvActuel)
				fmt.Print("|  3: ", Espace(20, tabSalle[i/9][j].monstre[2].Espece, pdv), "pv  |")
			} else if i%9 == 5 || i%9 == 6 || i%9 == 7 {
				fmt.Print("|                              |")
			} else if i%9 == 8 && tabSalle[i/9][j].bas {
				fmt.Print("|_____________    _____________|")
			} else if i%9 == 8 && !tabSalle[i/9][j].bas {
				fmt.Print("|______________________________|")
			}
		}
		println("")

	}
}

func (u User) AffichageInventaire() {
	fmt.Print("\033[H\033[2J")
	niveau := fmt.Sprintf("%d", u.Niveau)
	pdv := fmt.Sprintf("%d", u.PdvActuel)
	tour := fmt.Sprintf("%d", u.tour)
	TabStat := [][]string{
		{"------------------------------"},
		{Espace(29, "Nom Joueur", u.Nom)},
		{Espace(29, "Classe", u.Classe)},
		{Espace(29, "Niveau", niveau)},
		{Espace(29, "Points de vie ", pdv)},
		{Espace(29, "Nombre de tour", tour)},
		{"------------------------------"},
		{"                              "},
		{"                              "},
		{"goto direction                "},
		{"accessinventory               "},
		{"retour                        "},
		{"displayinfo                   "},
		{"quiter                        "},
		{"takepot                       "},
		{"poisonpot                     "},
		{"retour                        "},
		{"marchand                      "},
		{"forgeron                      "},
		{"skill coupdepoing lignemonstre"},
		{"pass                          "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
		{"                              "},
	}
	for i := 0; i < 36; i++ {
		for index := range TabStat[i] {
			fmt.Print("|", TabStat[i][index], "|")
		}
		if len(u.InventaireJoueur) > i {
			fmt.Print("	-", u.InventaireJoueur[i].NomObjet, " :", u.InventaireJoueur[i].Quantite)
		}
		fmt.Println()
	}
}

func (u User) displayInfoNain() {
	fmt.Print("\033[H\033[2J")
	niveau := fmt.Sprintf("%d", u.Niveau)
	pdv := fmt.Sprintf("%d", u.PdvActuel)
	tour := fmt.Sprintf("%d", u.tour)
	fmt.Println("------------------------------|", " ______   ___   _______  _______  ___      _______  __   __    ___   __    _  _______  _______ ")
	fmt.Println(Espace(28, "Nom Joueur", u.Nom), "|", "|      | |   | |       ||       ||   |    |   _   ||  | |  |  |   | |  |  | ||       ||       |")
	fmt.Println(Espace(28, "Classe", u.Classe), "|", "|  _    ||   | |  _____||    _  ||   |    |  |_|  ||  |_|  |  |   | |   |_| ||    ___||   _   |")
	fmt.Println(Espace(28, "Niveau", niveau), "|", "| | |   ||   | | |_____ |   |_| ||   |    |       ||       |  |   | |       ||   |___ |  | |  |")
	fmt.Println(Espace(30, "Points de vie ", pdv), "|", "| |_|   ||   | |_____  ||    ___||   |___ |       ||_     _|  |   | |  _    ||    ___||  |_|  |")
	fmt.Println(Espace(29, "Nombre de tour", tour), "|       ||   |  _____| ||   |    |       ||   _   |  |   |    |   | | | |   ||   |    |       |")
	fmt.Println("                              |", "|______| |___| |_______||___|    |_______||__| |__|  |___|    |___| |_|  |__||___|    |_______|")
	fmt.Println("                              |")
	fmt.Println("                              |")
	fmt.Println("                              |")
	fmt.Println("                              |", "                                       Nom du Joueur    : ", u.Nom)
	fmt.Println("                              |", "                                       Classe du Joueur : Nain")
	fmt.Println("                              |", "                                       Niveau           : ", string(u.Niveau+48))
	fmt.Println("                              |", "                                       Points de vie    : ", string(u.PdvActuel+48), "/", string(u.PdvMax))
	fmt.Println("                              |")
	fmt.Println("                              |                      Nom : ", u.Nom, "")
	fmt.Println("                              |                      Âge : 178 ans")
	fmt.Println("                              |                      Origine : Montagnes de Fer, royaume souterrain de Khar-Dûm")
	fmt.Println("                              |")
	fmt.Println("                              |                      Histoire de ", u.Nom, "")
	fmt.Println("                              |                      ", u.Nom, " est né dans le clan Forgeflamme, réputé pour ses artisans et forgerons d’exception. Dès son plus jeune âge, ")
	fmt.Println("                              |                      il a été initié au maniement du marteau et à la maîtrise du feu, héritage ancestral de son peuple. Les montagnes de Fer, ")
	fmt.Println("                              |                      leur refuge, sont un labyrinthe de cavernes riches en minerais précieux, mais aussi en dangers anciens.")
	fmt.Println("                              |")
	fmt.Println("                              |                      Le clan Forgeflamme n’est pas seulement connu pour ses forges : il est aussi le gardien d’un secret millénaire, ")
	fmt.Println("                              |                      une relique forgeée par les ancêtres, la Lame du Crépuscule, une épée légendaire capable de dompter le feu et la pierre. ")
	fmt.Println("                              |                      Cette lame a été perdue depuis plusieurs générations, emportée dans un effondrement mystérieux.")
	fmt.Println("                              |")
	fmt.Println("                              |                      ", u.Nom, ", animé par un sens profond de l’honneur et du devoir, a juré de retrouver la Lame du Crépuscule. ")
	fmt.Println("                              |                      Son voyage le mènera à travers des cavernes hantées, des royaumes de surface en guerre, et des alliances inattendues, ")
	fmt.Println("                              |                      avec pour seul guide la sagesse des runes et la force indomptable de son marteau.")
	fmt.Println("                              |")
	fmt.Println("                              |                      Mais plus que tout, ", u.Nom, " devra affronter une menace grandissante : ")
	fmt.Println("                              |                      une ancienne créature de pierre réveillée par la soif du pouvoir, menaçant de détruire ")
	fmt.Println("                              |                      tout ce que son peuple a construit.")
	fmt.Println("                              |")
	fmt.Println("                              |")
}

func (u User) displayInfoAssassin() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("------------------------------|", " ______   ___   _______  _______  ___      _______  __   __    ___   __    _  _______  _______ ")
	fmt.Println(Espace(28, "Nom Joueur", u.Nom), "|", "|      | |   | |       ||       ||   |    |   _   ||  | |  |  |   | |  |  | ||       ||       |")
	fmt.Println(Espace(28, "Classe", u.Classe), "|", "|  _    ||   | |  _____||    _  ||   |    |  |_|  ||  |_|  |  |   | |   |_| ||    ___||   _   |")
	fmt.Println(Espace(28, "Niveau", string(u.Niveau+48)), "|", "| | |   ||   | | |_____ |   |_| ||   |    |       ||       |  |   | |       ||   |___ |  | |  |")
	fmt.Println(Espace(30, "Points de vie ", string(u.PdvActuel+48)), "|", "| |_|   ||   | |_____  ||    ___||   |___ |       ||_     _|  |   | |  _    ||    ___||  |_|  |")
	fmt.Println("------------------------------|", "|       ||   |  _____| ||   |    |       ||   _   |  |   |    |   | | | |   ||   |    |       |")
	fmt.Println("                              |", "|______| |___| |_______||___|    |_______||__| |__|  |___|    |___| |_|  |__||___|    |_______|")
	fmt.Println("                              |")
	fmt.Println("                              |")
	fmt.Println("                              |")
	fmt.Println("                              |", "                                       Nom du Joueur    : ", u.Nom)
	fmt.Println("                              |", "                                       Classe du Joueur : Assassin")
	fmt.Println("                              |", "                                       Niveau           : ", string(u.Niveau+48))
	fmt.Println("                              |", "                                       Points de vie    : ", string(u.PdvActuel+48), "/", string(u.PdvMax))
	fmt.Println("                              |")
	fmt.Println("                              |                      Nom : ", u.Nom, "")
	fmt.Println("                              |                      Âge : 29 ans")
	fmt.Println("                              |                      Origine : La cité-État d’Eryndal, ruelles obscures et marchés noirs")
	fmt.Println("                              |")
	fmt.Println("                              |                      Histoire de ", u.Nom, "")
	fmt.Println("                              |                      ", u.Nom, " est né dans les bas-fonds d’Eryndal, une cité rongée par la corruption et les complots. Orphelin très jeune, ")
	fmt.Println("                              |                      il a survécu grâce à son intelligence, sa discrétion, et surtout à ses talents innés pour se fondre dans l’ombre. ")
	fmt.Println("                              |                      Recruté par la guilde secrète des Dagues du Crépuscule, il a été formé aux arts du silence, du poison, et de la disparition.")
	fmt.Println("                              |")
	fmt.Println("                              |                      Son passé est un mystère même pour ses alliés : certains disent qu’il serait le fils illégitime d’un noble déchu, ")
	fmt.Println("                              |                      d’autres qu’il a un pacte ancien avec une entité des ombres, ce qui expliquerait ses réflexes quasi surhumains.")
	fmt.Println("                              |                      ")
	fmt.Println("                              |                      ", u.Nom, " travaille dans l’ombre pour des causes variées — assassinats politiques, élimination de traîtres, ")
	fmt.Println("                              |                      ou missions de sabotage — mais il suit un code strict : jamais d’innocents, jamais pour de l’argent seul. Sa vraie motivation ? ")
	fmt.Println("                              |                      Démasquer la conspiration qui a détruit sa famille et dévoré la ville de l’intérieur.")
	fmt.Println("                              |                      ")
	fmt.Println("                              |                      Chaque cible éliminée est une pièce du puzzle qui le rapproche de la vérité, mais aussi des ennemis plus puissants, ")
	fmt.Println("                              |                      tapis dans l’ombre, prêts à tout pour préserver leurs secrets.")
	fmt.Println("                              |                      ")
	fmt.Println("                              |                      ")
	fmt.Println("                              |                      ")
	fmt.Println("                              |")
}

func (u User) displayInfoElfe() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("------------------------------|", " ______   ___   _______  _______  ___      _______  __   __    ___   __    _  _______  _______ ")
	fmt.Println(Espace(28, "Nom Joueur", u.Nom), "|", "|      | |   | |       ||       ||   |    |   _   ||  | |  |  |   | |  |  | ||       ||       |")
	fmt.Println(Espace(28, "Classe", u.Classe), "|", "|  _    ||   | |  _____||    _  ||   |    |  |_|  ||  |_|  |  |   | |   |_| ||    ___||   _   |")
	fmt.Println(Espace(28, "Niveau", string(u.Niveau+48)), "|", "| | |   ||   | | |_____ |   |_| ||   |    |       ||       |  |   | |       ||   |___ |  | |  |")
	fmt.Println(Espace(30, "Points de vie ", string(u.PdvActuel+48)), "|", "| |_|   ||   | |_____  ||    ___||   |___ |       ||_     _|  |   | |  _    ||    ___||  |_|  |")
	fmt.Println("------------------------------|", "|       ||   |  _____| ||   |    |       ||   _   |  |   |    |   | | | |   ||   |    |       |")
	fmt.Println("                              |", "|______| |___| |_______||___|    |_______||__| |__|  |___|    |___| |_|  |__||___|    |_______|")
	fmt.Println("                              |")
	fmt.Println("                              |")
	fmt.Println("                              |")
	fmt.Println("                              |", "                                       Nom du Joueur    : ", u.Nom)
	fmt.Println("                              |", "                                       Classe du Joueur : Elfe")
	fmt.Println("                              |", "                                       Niveau           : ", string(u.Niveau+48))
	fmt.Println("                              |", "                                       Points de vie    : ", string(u.PdvActuel+48), "/", string(u.PdvMax))
	fmt.Println("                              |")
	fmt.Println("                              |                      Nom : ", u.Nom, "")
	fmt.Println("                              |                      Âge : 243 ans (jeune adulte pour une elfe)")
	fmt.Println("                              |                      Origine : Forêt d’Élyndra, royaume sylvestre des Hauts-Arbres")
	fmt.Println("                              |")
	fmt.Println("                              |                      Histoire d’", u.Nom, "")
	fmt.Println("                              |                      ", u.Nom, " est née sous la canopée millénaire d’Élyndra, un sanctuaire naturel préservé des hommes et des conflits, ")
	fmt.Println("                              |                      où la magie coule dans chaque feuille et chaque ruisseau. Issue d’une lignée de gardiennes, ")
	fmt.Println("                              |                      elle est liée dès son enfance aux esprits de la forêt et aux anciennes légendes elfiques.")
	fmt.Println("                              |")
	fmt.Println("                              |                      Dotée d’une sensibilité unique, ", u.Nom, " communique avec les animaux et peut manipuler doucement les vents ")
	fmt.Println("                              |                      pour guider les voyageurs ou détourner les ennemis. Sa nature pacifique cache pourtant une détermination farouche : ")
	fmt.Println("                              |                      elle veille à ce que la forêt ne soit jamais souillée par la guerre ni par la cupidité des humains.")
	fmt.Println("                              |                      ")
	fmt.Println("                              |                      Mais une ombre grandissante menace Élyndra. Des forces obscures venues des terres dévastées cherchent à corrompre ")
	fmt.Println("                              |                      la magie de la forêt pour leurs propres desseins. ", u.Nom, " a été choisie par les Anciens pour partir en quête des Cristaux de Lune, ")
	fmt.Println("                              |                      des artefacts capables de purifier la terre et de restaurer l’équilibre naturel.")
	fmt.Println("                              |                      ")
	fmt.Println("                              |                      Son voyage la confrontera aux conflits entre races, aux choix moraux complexes, et aux secrets enfouis dans les racines mêmes de la forêt. ")
	fmt.Println("                              |                      Entre tradition et changement, ", u.Nom, " devra définir ce que signifie vraiment être une gardienne.")
	fmt.Println("                              |                      ")
	fmt.Println("                              |                      ")
	fmt.Println("                              |")
}
