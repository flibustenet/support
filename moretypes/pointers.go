//go:build OMIT

package main

import "fmt"

func valeur(j string) {
	j = "valeur"
}

func pointeur(j *string) {
	*j = "pointeur"
}

func main() {
	i := "i"
	fmt.Println("i=", i)
	valeur(i)
	fmt.Println("valeur i=", i)
	var p *string = &i
	pointeur(p)
	fmt.Println("pointeur i=", i)
}
