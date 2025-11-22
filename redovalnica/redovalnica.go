// Package redovalnica provides basic functionality for managing student grades.
//
// The package allows adding grades to students, printing all grades, and evaluating
// final success based on configurable parameters such as minimum number of grades
// and valid grade range.
//
// Example usage:
//
//	studenti := map[string]redovalnica.Student{
//		"63210001": {Ime: "Ana", Priimek: "Novak"},
//		"63210002": {Ime: "Boris", Priimek: "Kralj"},
//	}
//
//	redovalnica.DodajOceno(studenti, "63210001", 9)
//	redovalnica.IzpisVsehOcen(studenti)
//	redovalnica.IzpisiKoncniUspeh(studenti, 2, 1, 10)
//
// The final evaluation is based on the following logic:
//   - If the average grade is 9.0 or more: "Odličen študent!"
//   - If the average grade is between 6.0 and 8.99: "Povprečen študent"
//   - If the average grade is below 6.0: "Neuspešen študent."
//   - If the number of grades is below the minimum: "premalo ocen"
//   - If any grade is outside the allowed range: "neveljavne ocene"
package redovalnica

import "fmt"


type Student struct {
	ime  string
	priimek string
	ocene []int
}


func DodajOceno(studenti map[string]Student, vpisna string, ocena int) {
	student := studenti[vpisna]
	student.Ocene = append(student.Ocene, ocena)
	studenti[vpisna] = student
}

func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, s := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, s.Ime, s.Priimek, s.Ocene)
	}
}

func IzpisiKoncniUspeh(studenti map[string]Student, stOcen, minOcena, maxOcena int) {
	for vpisna, s := range studenti {
		povp := povprecje(studenti, vpisna)
		if len(s.Ocene) < stOcen {
			fmt.Printf("%s %s: premalo ocen\n", s.Ime, s.Priimek)
			continue
		}
		valid := true
		for _, o := range s.Ocene {
			if o < minOcena || o > maxOcena {
				valid = false
				break
			}
		}
		if !valid {
			fmt.Printf("%s %s: neveljavne ocene\n", s.Ime, s.Priimek)
			continue
		}

		var rezultat string
		switch {
		case povp >= 9.0:
			rezultat = "Odličen študent!"
		case povp >= 6.0:
			rezultat = "Povprečen študent"
		default:
			rezultat = "Neuspešen študent."
		}
		fmt.Printf("%s %s: povprečna ocena %.2f -> %s\n", s.Ime, s.Priimek, povp, rezultat)
	}
}

func povprecje(studenti map[string]Student, vpisna string) float64 {
	student, ok := studenti[vpisna]
	if !ok || len(student.Ocene) == 0 {
		return -1.0
	}
	sum := 0
	for _, o := range student.Ocene {
		sum += o
	}
	return float64(sum) / float64(len(student.Ocene))
}
