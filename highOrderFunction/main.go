package main

import "fmt"

type student struct {
	name       string
	grade      int
	attendance int
}

func filter(students []student, f func(student) bool) []student {
	var arr []student
	for _, s := range students {
		if f(s) {
			arr = append(arr, s)
		}
	}
	return arr
}

func genFilterFunc(schoolName string) func(student) bool {
	if schoolName == "A" {
		return func(s student) bool {
			if s.grade >= 5 {
				return true
			}
			return false
		}
	} else if schoolName == "B" {
		return func(s student) bool {
			if s.grade >= 5 && s.attendance >= 7 {
				return true
			}
			return false
		}

	} else {
		return func(s student) bool { return true }
	}
}

func main() {
	s1 := student{
		name:       "Bob",
		grade:      4,
		attendance: 8,
	}
	s2 := student{
		name:       "Alice",
		grade:      8,
		attendance: 4,
	}
	s := []student{s1, s2}
	f := filter(s, func(s student) bool {
		if s.grade >= 5 {
			return true
		}
		return false
	})
	fmt.Printf("Student filter simple: %v\n", f)

	// Génération d'une fonction en type de retour

	filterFunc := genFilterFunc("Z")
	f = filter(s, filterFunc)
	fmt.Printf("Student filter school Z: %v\n", f)

	// Génération d'une fonction pour l'école A
	filterFunc = genFilterFunc("A")
	f = filter(s, filterFunc)
	fmt.Printf("Student filter school A: %v\n", f)

	// Génération d'une fonction pour l'école B
	filterFunc = genFilterFunc("B")
	f = filter(s, filterFunc)
	fmt.Printf("Student filter school B: %v\n", f)

}
