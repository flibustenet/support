//go:build OMIT

package main

import "fmt"

func main() {
	// Initialisation sans valeurs (chaines vides par défaut)
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// initialisation avec valeurs
	// bien spécifier le nombre de valeurs 6 pour
	// différencier d'un slice à nombre de valeurs variables
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// si moins de valeurs, le reste est initialisé à zéro
	prz := [6]int{2, 3, 5}
	fmt.Println(prz)

	// initialisation avec taille automatique (mais toujours fixe)
	pr := [...]int{1, 2, 3, 4}
	fmt.Println(pr)
}
