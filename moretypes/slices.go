//go:build OMIT

package main

import "fmt"

func main() {
	// Tableau de taille fixe
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes, len(primes))

	// Slice, tableau de taille variable
	// ici initialisé à partir du tableau précédent
	var s []int = primes[1:4]
	fmt.Println(s, len(s))

	// ici initialisé avec des valeurs et une taille automatique
	// du nombre de valeurs
	x := []int{1, 2, 3, 4}
	fmt.Println(x, len(x))
}
