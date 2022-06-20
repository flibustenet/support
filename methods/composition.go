//go:build OMIT
package main

import (
	"fmt"
)

type Auteur struct {
	nom    string
	prenom string
}

func (a *Auteur) NomPrenom() string {
	return a.nom + " " + a.prenom
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
	return fmt.Sprintf("Titre = %s-%s Auteur = %s", l.titre, l.tome, l.NomPrenom())
}

func (l *Livre) String2() string {
	return fmt.Sprintf("Titre = %s-%s Auteur = %s %s", l.titre, l.tome, l.nom, l.prenom)
}

func (l *Livre) String3() string {
	return fmt.Sprintf("Titre = %s-%s Auteur = %s", l.titre, l.tome, l.Auteur.String())
}

func main() {
	livre := &Livre{
		Auteur: Auteur{"Hugo", "Victor"},
		titre:  "Les misérables",
		tome:   "Tome 1",
	}
	fmt.Println(livre.String())
	fmt.Println(livre.String2())
	fmt.Println(livre.String3())
}
