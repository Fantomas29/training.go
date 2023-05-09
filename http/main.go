package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	//ecriture d'un message en reponse HTTP
	fmt.Printf("%v %v\n", r.Method, r.URL)
	fmt.Fprintf(w, "Hello gophers!")
}

// une requete ressemble a cela
// http://localhost:9000/search?t=go&p=1
func search(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("t")
	p := r.URL.Query().Get("p")
	fmt.Fprintf(w, "Searching for term=%v. Current page=%v", t, p)
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	//ecriture d'un message en reponse HTTP
	fmt.Printf("%v %v\n", r.Method, r.URL)
	fmt.Fprintf(w, "Goodbye gophers!")
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "login.html")
		return
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() failed. err=%v", err)
			return
		}
		fmt.Fprintf(w, "Go login POST. value=%v\n", r.PostForm)
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username == "Go" && password == "rocks" {
			fmt.Fprintf(w, "You are now logged\n")
		} else {
			fmt.Fprintf(w, "Wrong username or password\n")
		}

	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/goodbye", goodbye)
	http.HandleFunc("/search", search)
	http.HandleFunc("/login", login)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
