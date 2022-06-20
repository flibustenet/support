//go:build OMIT

package main

import (
	"fmt"
)

// Exercice : ajouter la couleur à la voiture et trier les véhicules sur la couleur

type Vehicule interface {
	fmt.Stringer
	NbRoues() int
}

type Velo struct {
	Couleur string
}

func (v *Velo) NbRoues() int {
	return 2
}
func (v *Velo) String() string {
	return fmt.Sprintf("Vélo %s", v.Couleur)
}

type Voiture struct {
	Carburant string
}

func (v *Voiture) NbRoues() int {
	return 4
}
func (v *Voiture) String() string {
	return fmt.Sprintf("Voiture %s", v.Carburant)
}

func main() {
	vehicules := []Vehicule{
		&Velo{Couleur: "rouge"},
		&Voiture{Carburant: "essence"},
	}
	for _, v := range vehicules {
		fmt.Println(v)
	}
}
