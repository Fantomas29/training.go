package fonctions

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// src: nom de fichier source
//
// old/new: ancien mot/nouveau mot
//
// occ : nombre d'occurences de old
//
// lines: slice des numéros de lignes où old a été trouvé
//
// err : erreur de la fonction
func FindReplaceFile(src, old, new, dest string) (occ int, lines []int, nbLines int, err error) {

	//test de lecture du fichier src
	dat, err := ioutil.ReadFile(src)
	if err != nil {
		return 0, lines, 0, err
	}

	//si le fichier src est vide
	if len(dat) == 0 {
		return occ, lines, 0, errors.New("Empty content")
	}

	//ouverture du fichier source
	fileSrc, err := os.Open(src)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier: ", err)
		return
	}

	//creation d'un nouveau fichier destination à partir du chemin dest
	fileDest, err := os.Create(dest)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier:", err)
		return
	}
	//fermeture des fichiers
	defer fileDest.Close()
	defer fileSrc.Close()

	//creation des variables de retour
	var (
		trouve      bool
		nbTrouve    int
		ligneRetour string
		i           int
	)

	//scan le fichier src lignes par lignes
	scanner := bufio.NewScanner(fileSrc)
	for scanner.Scan() {
		line := scanner.Text()
		//appel de la fonction qui modifie les lignes
		trouve, ligneRetour, nbTrouve = ProcessLine(line, old, new)
		occ += nbTrouve
		//si on a trouve une ligne a remplace on ajoute le numero de la ligne dans le slice
		if trouve {
			lines = append(lines, i)
		}
		i++
		//ecriture de la nouvelle ligne dans le fichier destination
		writer := bufio.NewWriter(fileDest)
		fmt.Fprintln(writer, ligneRetour)
		writer.Flush()
	}
	return occ, lines, i, nil
}
