package hangman

import "fmt"

func DrawWelcome() {
	fmt.Println(` _                                             
	| |                                            
	| |__   __ _ _ __   __ _ _ __ ___   __ _ _ __  
	|  _ \ / _  |  _ \ / _  |  _   _ \ / _  |  _ \ 
	| | | | (_| | | | | (_| | | | | | | (_| | | | |
	|_| |_|\__ _|_| |_|\__  |_| |_| |_|\__ _|_| |_|
			    _/  |                      
		   	   |___/                                                
	`)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
		+---+
		|   |
		O   |
	       /|\  |
	       / \  |
		    |
	  =========
		`
	case 1:
		draw = `
		+---+
		|   |
		O   |
	       /|\  |
	       /    |
		    |
	  =========
		`
	case 2:
		draw = `
		+---+
		|   |
		O   |
	       /|\  |
		    |
		    |
	  =========
		`
	case 3:
		draw = `
		+---+
		|   |
		O   |
	       /|   |
		    |
		    |
	  =========
		`
	case 4:
		draw = `
		+---+
		|   |
		O   |
		|   |
		    |
		    |
	  =========
		`
	case 5:
		draw = `
		+---+
		|   |
		O   |
		    |
		    |
		    |
	  =========
		`
	case 6:
		draw = `
		+---+
		|   |
		    |
		    |
		    |
		    |
	  =========
		`
	case 7:
		draw = `
		
		`
	}
	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed : ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used : ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Good guess !\n")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' already used\n", guess)
	case "badGuess":
		fmt.Printf("Bad guess, '%s' is not in the word\n", guess)
	case "lost":
		fmt.Print("You lost :(! The word was : ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("YOU WON ! the word was : ")
		drawLetters(g.Letters)
	}

}

func drawLetters(l []string) {
	for _, c := range l {

		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
