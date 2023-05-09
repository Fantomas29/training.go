package main

import (
	"fmt"

	"training.go/goplace/fonctions"
)

func main() {

	var (
		source, old, new, dest string
		occurence, nbLigne     int
		lignes                 []int
		erreur                 error
	)
	source = "fonctions/fichierTest.txt"
	old = "pom"
	new = "apple"
	dest = "resultat.txt"
	occurence, lignes, nbLigne, erreur = fonctions.FindReplaceFile(source, old, new, dest)

	if erreur != nil {
		println("erreur lors de l'appel à la fonction : ", erreur)
		return
	}

	fmt.Printf("Nombre d'occurences de %s = %d\nnombre de lignes : %d\nsur les lignes : %v\n", old, occurence, nbLigne, lignes)
	fmt.Print("lignes : [ ")
	len := len(lignes)
	for i, l := range lignes {
		fmt.Printf("%v", l)
		if i < len-1 {
			fmt.Printf(" - ")
		}
	}
	fmt.Println(" ] ")
}
