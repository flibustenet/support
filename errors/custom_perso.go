//go:build OMIT

package main

import (
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

func main() {
	res, err := divCustom(4, 0)
	if e, ok := err.(*MyError); ok {
		fmt.Printf("Erreur prévue : msg = %s", e.msg)
		return
	}
	if err != nil {
		fmt.Printf("Retour : %v\n", err)
	}
	fmt.Printf("Résultat : %.2f\n", res)
}
