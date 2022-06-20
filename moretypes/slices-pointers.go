//go:build OMIT

package main

import "fmt"

func main() {
	// tableau de taille fixe
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("names=", names)

	// slices du tableau précédent
	a := names[0:2]
	b := names[1:3]
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	// ici essayez de modifier un autre élément du tableau
	// à partir de names ou a
	b[0] = "XXX"

	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("names=", names)
}
