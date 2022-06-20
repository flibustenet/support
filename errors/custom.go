//go:build OMIT

package main

import (
	"errors"
	"fmt"
)

// valeur d'erreur prédéfinie
var monErreur = errors.New("Division par zéro")

// renvoi éventuellement l'erreur prédéfinie
func div(i, j float64) (float64, error) {
	if j == 0 {
		return 0, monErreur
	}
	return i / j, nil
}

func main() {
	res, err := div(4, 0)
	if err != nil {
		fmt.Printf("Retour : %v\n", err)
	}
	fmt.Printf("Résultat : %.2f\n", res)
}
