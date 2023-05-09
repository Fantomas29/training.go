package main

import (
	"fmt"
)

func main() {

	m := map[string]bool{
		"Beatles":        true,
		"Rolling Stones": true,
		"Led Zeppelin":   false,
	}
	for band, present := range m {
		if band == "Rolling Stones" {
			delete(m, band)
		}
		fmt.Printf("%v=%v\n", band, present)
		fmt.Println("boucle for : ", m)
	}

	fmt.Println("apres la boucle : ", m)
}
