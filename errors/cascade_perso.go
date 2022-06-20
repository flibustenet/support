//go:build OMIT

package main

import (
	"errors"
	"fmt"
)

// type d'erreur personalisée
type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

// renvoi éventuellement l'erreur personalisée
func divCustom(i, j float64) (float64, error) {
	if j == 0 {
		return 0, &MyError{fmt.Sprintf("%.2f/0 n'est pas possible", i)}
	}
	return i / j, nil
}

// appel intermédiaire avec masquage de l'erreur originale
func appel(i, j float64) (float64, error) {
	res, err := divCustom(i, j)
	if err != nil {
		return 0, fmt.Errorf("Appel div avec %.2f et %.2f : %v", i, j, err) // avec %v
	}
	return res, nil
}

// appel intermédiaire sans masquage de l'erreur originale
func appelWrap(i, j float64) (float64, error) {
	res, err := divCustom(i, j)
	if err != nil {
		return 0, fmt.Errorf("Appel div avec %.2f et %.2f : %w", i, j, err) // avec %w
	}
	return res, nil
}

func main() {
	// appel intermédiaire
	res, err := appel(4, 0) // changer avec appelWrap
	if err != nil {
		e := &MyError{}
		if errors.As(err, &e) {
			fmt.Printf("Retour MyError : %s\n", e.msg)
		} else {
			fmt.Printf("Retour : %v\n", err)
		}
	}
	fmt.Printf("Résultat : %.2f\n", res)

}
