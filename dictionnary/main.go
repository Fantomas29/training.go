package main

import (
	"flag"
	"fmt"
	"os"

	"training.go/dictionnary/dictionnary"
)

func main() {
	action := flag.String("action", "list", "Action to perform on the dictionnary")

	d, err := dictionnary.New("./badger")
	handleErr(err)
	defer d.Close()

	flag.Parse()
	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	default:
		fmt.Printf("Unknown action : %v\n", *action)
	}
}

func actionList(d *dictionnary.Dictionnary) {
	words, entries, err := d.List()
	handleErr(err)
	fmt.Println("dictionnary content")
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func actionAdd(d *dictionnary.Dictionnary, args []string) {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleErr(err)
	fmt.Printf("'%v' added to the dictionnary\n", word)
}

func actionDefine(d *dictionnary.Dictionnary, args []string) {
	word := args[0]
	entry, err := d.Get(word)
	handleErr(err)
	fmt.Println(entry)
}

func actionRemove(d *dictionnary.Dictionnary, args []string) {
	word := args[0]
	err := d.Remove(word)
	handleErr(err)
	fmt.Printf("'%v' word has been removed\n", word)
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionnary error : %v\n", err)
		os.Exit(1)
	}
}
