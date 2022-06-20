//go:build OMIT

package main

import "fmt"

// Index renvoi l'index de x dans s ou -1 s'il n'y est pas
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v et x sont de type T, qui doivent être "comparable"
		// nous pouvons donc utiliser == (et !=)
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index marche avec int
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index marche aussi avec string
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
