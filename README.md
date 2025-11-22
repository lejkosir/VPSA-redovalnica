## Package redovalnica provides basic functionality for managing student grades.

The package allows adding grades to students, printing all grades, and evaluating
final success based on configurable parameters such as minimum number of grades
and valid grade range.

Example usage:
```go
	studenti := map[string]redovalnica.Student{
		"63210001": {Ime: "Ana", Priimek: "Novak"},
		"63210002": {Ime: "Boris", Priimek: "Kralj"},
	}

	redovalnica.DodajOceno(studenti, "63210001", 9)
	redovalnica.IzpisVsehOcen(studenti)
	redovalnica.IzpisiKoncniUspeh(studenti, 2, 1, 10)
```
The final evaluation is based on the following logic:
  - If the average grade is 9.0 or more: "Odličen študent!"
  - If the average grade is between 6.0 and 8.99: "Povprečen študent"
  - If the average grade is below 6.0: "Neuspešen študent."
  - If the number of grades is below the minimum: "premalo ocen"
  - If any grade is outside the allowed range: "neveljavne ocene"
