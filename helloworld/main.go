package notmain

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func switchWord(src, old, new string) (occ int, lines []int, err error) {

	dat, err := ioutil.ReadFile(src)
	if err != nil {
		return 0, lines, err
	}

	if len(dat) == 0 {
		return occ, lines, errors.New("Empty content")
	}
	return occ, lines, nil
}

func main() {

	names := []string{"Toto", "Bobette", "Johan", "John"}

	for _, name := range names {
		fmt.Printf("Hello %s\n", name)
		defer fmt.Printf("goodbye %s\n", name)
	}
}
