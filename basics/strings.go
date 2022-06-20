// go:build OMIT

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Leila" // essayer de mettre un tréma sur le i

	fmt.Printf("len = %d unicode = %d\n", len(s), utf8.RuneCountInString(s))

	asBytes := []byte(s)
	fmt.Printf("len bytes = %d\n", len(asBytes))

	asRunes := []rune(s)
	fmt.Printf("len runes = %d\n", len(asRunes))

}
