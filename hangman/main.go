package main

import (
	"fmt"
	"os"

	"training.go/hangman/dictionnary"
	"training.go/hangman/hangman"
)

func main() {
	err := dictionnary.Load("dictionnaire.txt")
	if err != nil {
		fmt.Printf("Could not load dictionnary : %v\n", err)
		os.Exit(1)
	}
	g, err := hangman.New(7, dictionnary.PickWord())
	hangman.DrawWelcome()
	guess := ""
	for {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			os.Exit(0)

		}
		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal : %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)

	}

}
