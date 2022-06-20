//go:build OMIT

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// Exercice : Ajouter un Reader maison qui lit x chars donnés en paramètre, maximum 26

var lettres = []byte("abcdefghijklmnopqrstuvwxyz")

func main() {
	for i, r := range []io.Reader{
		strings.NewReader("String"),
		bytes.NewBufferString("Byte buffer"),
		// Ici
	} {
		data, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatalf("Impossible de lire %d", i)
		}
		fmt.Printf("%d: %s\n", i, data)
	}
}
