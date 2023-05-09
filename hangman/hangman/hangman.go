package hangman

import (
	"fmt"
	"strings"
)

type Game struct {
	State        string   //etat du jeu
	Letters      []string //Lettre dans le mot a trouve
	FoundLetters []string //Lettres trouvees
	UsedLetters  []string //Lettres utilisees
	TurnsLeft    int      //Nombres de tours restant
}

func New(turns int, word string) (*Game, error) {
	if len(word) < 3 {
		return nil, fmt.Errorf("Word '%s' must be at least 3 characters. got =%v.", word, len(word))
	}
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}
	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}
	return g, nil
}

func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)
	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)

		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	}
	if !letterInWord(guess, g.Letters) {
		g.State = "badGuess"
		g.UsedLetters = append(g.UsedLetters, guess)
		g.TurnsLeft--
		if g.TurnsLeft == 0 {
			g.State = "lost"
		}

	}

}

func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}

func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}

func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}
