package main

type Monstre struct {
	Espece    string
	PdvMax    int
	PdvActuel int
}

func initMonstre(espece string) Monstre {
	switch espece {
	case "orc":
		return Monstre{"orc", 5, 5}
	case "gobelin":
		return Monstre{"gobelin", 9, 9}
	default:
		return Monstre{"elfs", 2, 2}
	}
}

func ActionMonstre(tabSalle [][]Salles) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			for index := range tabSalle[i][j] {

			}
		}
	}
}
