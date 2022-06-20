// go:build OMIT

package main

import "fmt"

// modifier add pour convertir les string en nombre
// renvoyer le résultat sous forme de string
func add(xs ...string) string {
	res := 0
	for _, x := range xs {
		fmt.Println("Add ", x)
		res += x
	}
	return res
}

func main() {
	s := "42, 13, 8, 3, 5"
	// utiliser strings.Fields
}
