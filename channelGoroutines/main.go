package main

import (
	"fmt"
)

func main() {
	inputChan := make(chan string)

	// Demande à l'utilisateur s'il est disponible ce midi
	go func() {
		for {
			var input string
			fmt.Println("Es-tu disponible ce midi ? (oui/non)")
			fmt.Scanln(&input)
			inputChan <- input
		}
	}()

	// Boucle tant que l'utilisateur ne répond pas "oui"
	for {
		input := <-inputChan
		if input == "oui" {
			fmt.Println("Super, à tout à l'heure !")
			close(inputChan)
			break
		} else if input == "non" {
			fmt.Println("Comment ça non ? je suis triste :'(, réponse impossible")
		} else {
			fmt.Println("Réponse impossible, choisie une autre option.")
		}
	}
}
