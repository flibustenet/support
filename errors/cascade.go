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

// appel intermédiaire avec masquage de l'erreur originale
func appel(i, j float64) (float64, error) {
	res, err := div(i, j)
	if err != nil {
		return 0, fmt.Errorf("Appel div avec %.2f et %.2f : %v", i, j, err) // avec %v
	}
	return res, nil
}

// appel intermédiaire sans masquage de l'erreur originale
func appelWrap(i, j float64) (float64, error) {
	res, err := div(i, j)
	if err != nil {
		return 0, fmt.Errorf("Appel div avec %.2f et %.2f : %w", i, j, err) // avec %w
	}
	return res, nil
}

func main() {
	res, err := div(4, 0) // appel direct
	if err != nil {
		if err == monErreur { // comparaison simple et directe, sans cascade
			fmt.Printf("Retour monErreur : %v\n", err)
		} else {
			fmt.Printf("Retour : %v\n", err)
		}
	}
	fmt.Printf("Résultat : %.2f\n", res)

	// appel intermédiaire
	res, err = appel(4, 0) // changer avec appelWrap
	if err != nil {
		if errors.Is(err, monErreur) {
			fmt.Printf("Retour monErreur : %v\n", err)
		} else {
			fmt.Printf("Retour : %v\n", err)
		}
	}
	fmt.Printf("Résultat : %.2f\n", res)

}
