//go:build OMIT
package main

import (
	"fmt"
)

type Auteur struct {
	nom    string
	prenom string
}

func (a *Auteur) concat(x, y string) string {
	return x + " " + y
}
func (a *Auteur) NomPrenom() string {
	return a.concat(a.nom, a.prenom)
}

func (a *Auteur) String() string {
	return a.NomPrenom()
}

type Livre struct {
	Auteur
	titre string
	tome  string
}

func (l *Livre) String() string {
	return fmt.Sprintf("Titre = %s\nAuteur = %s", l.concat(l.titre, l.tome), l.NomPrenom())
}

func (l *Livre) concat(x, y string) string {
	return x + "-" + y
}

func main() {
	livre := &Livre{
		Auteur: Auteur{"Hugo", "Victor"},
		titre:  "Les misérables",
		tome:   "Tome 1",
	}
	fmt.Println(livre)
}
