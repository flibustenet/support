// go:build OMIT

package main

import (
	"fmt"
	"strings"
)

//
// Les variables et les fonctions
/////////////////////////////////

// Déclarations globales

// exportées à l'aide d'une majuscule
// pourra être utilisé à l'aide de nom_du_package.Glob
var Glob int
var Flt float64

// locales au package
var loc int

// initialisées avec inférence de type
var ini = " -> "

// Déclaration du type de variable à passer en paramètre
// et en retour de fonction
// Fonction avec une majuscule, exportée nom_du_package.Multi
func Multi(nb int, ch string) string {
	// à n'importe quel endroit du package on peut utiliser les variables globales
	// sauf si elles ont été masquées par une variable locale du même nom

	// variables locales
	var i int       // == 0 par défaut
	var s string    // == "" par défaut
	var u = "texte" // inférence de type

	// OU

	// raccourcis initialisé avec inférence de type
	res := ini // utilise la valeur de la variable globale ini

	i = 99

	for i := 0; i < nb; i++ {
		// `i` masque le `i` précédent 99 (shadowing)
		// ici `maj` est locale à la boucle
		maj := strings.ToUpper(ch)
		res += maj
	}
	// impossible d'utiliser `maj` ici
	fmt.Println("i = ", i) // i == 99 n'a pas été modifié par la boucle

	// Le programme ne va pas compiler si les variables locales
	// ne sont pas utilisées.
	// error : declared but not used

	// utilisation de _ pour tricher
	_ = s
	_ = u

	return res
}

// les variables nommées n'ont pas besoin d'être initialisées
// ni retournées explicitement
// peu utilisé sauf sur petites fonctions
//
// fonction utilisable uniquement au sein du package
// éventuellement dans un autre fichier de code du même package
func named() (x int, y int) {
	x = 5
	y = 3
	return
}

func main() {
	res := Multi(4, "ok")
	fmt.Println("res = ", res)

	x, y := named()
	fmt.Println("Named return", x, y)
}
