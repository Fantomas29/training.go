package hangman

import "testing"

func TestLetterInWOrld(t *testing.T) {
	word := []string{"t", "o", "t", "o"}
	guess := "t"
	hasLetter := letterInWord(guess, word)
	if !hasLetter {
		t.Errorf("Mot %s contient la lettre %s et j'ai reçu=%v", word, guess, hasLetter)
	}
}

func TestLetterNotInWOrld(t *testing.T) {
	word := []string{"t", "o", "t", "o"}
	guess := "x"
	hasLetter := letterInWord(guess, word)
	if hasLetter {
		t.Errorf("Mot %s ne contient pas la lettre %s et j'ai reçu=%v", word, guess, hasLetter)
	}
}

func TestInvalidWord(t *testing.T) {
	_, err := New(3, "")

	if err == nil {
		t.Errorf("Error should be returned when using an invalid word =''")
	}

}

func TestGameGoodGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	validState(t, "goodGuess", g.State)
}

func TestGameBadGuess(t *testing.T) {
	g, _ := New(4, "gigolo")
	g.MakeAGuess("x")
	validState(t, "badGuess", g.State)
}

func TestGameAlreadyGuessed(t *testing.T) {
	g, _ := New(6, "used")
	g.MakeAGuess("u")
	g.MakeAGuess("u")
	validState(t, "alreadyGuessed", g.State)
}

func TestGameWon(t *testing.T) {
	g, _ := New(6, "win")
	g.MakeAGuess("w")
	g.MakeAGuess("i")
	g.MakeAGuess("n")
	validState(t, "won", g.State)
}

func TestGameLost(t *testing.T) {
	g, _ := New(4, "eaten")
	g.MakeAGuess("r")
	g.MakeAGuess("u")
	g.MakeAGuess("p")
	g.MakeAGuess("m")
	validState(t, "lost", g.State)
}

func validState(t *testing.T, expectedState, actualState string) bool {
	if expectedState != actualState {
		t.Errorf("State should be '%v'. got '%v'", expectedState, actualState)
		return false
	}
	return true
}
