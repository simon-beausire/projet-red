package main

func (u *User) poposoin() {
	valeur := "potion_de_soin"
	for i, v := range u.InventaireJoueur {
		if v.NomObjet == valeur {
			u.ajoutervie(3)
			u.InventaireJoueur = append(u.InventaireJoueur[:i], u.InventaireJoueur[i+1:]...)
		}
	}
}
