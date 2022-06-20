//go:build OMIT

package main

import (
	"errors"
	"fmt"
)

func previsible(i, j float64) (float64, error) {
	if j == 0 {
		return 0, errors.New("Pas de division par zéro")
	}
	return i / j, nil
}

func imprevisible(i int) string {
	return []string{"zero", "un", "deux"}[i]
}

func main() {
	res, err := previsible(4, 0)
	if err != nil {
		fmt.Printf("Erreur prévisible : %v\n", err)
	}
	fmt.Printf("res=%.2f\n", res)

	imprevisible(5) // runtime error: index out of range [5] with length 3
}
