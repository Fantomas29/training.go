package fonctions

import (
	"strings"
)

// line: ligne a traiter
//
// old/new: ancien mot/nouveau mot
//
// found: vrai si au moins une occurence trouvée
//
// res: resultat de remplacement (res == line si aucun changement !)
//
// occ: nombre d'occurences de old dans la ligne
func ProcessLine(line, old, new string) (found bool, res string, occ int) {

	lineUp := strings.ToUpper(line)
	oldUp := strings.ToUpper(old)

	//si on a au moins une occurence du mot contenu dans old
	if strings.Contains(lineUp, oldUp) {
		found = true
		//on compte son nombre d'occurences
		occ = strings.Count(lineUp, oldUp)
		//on remplace old par new et on met le resultat dans res
		res = strings.Replace(line, old, new, -1)
		return found, res, occ
	} else {
		return found, line, occ
	}

}
