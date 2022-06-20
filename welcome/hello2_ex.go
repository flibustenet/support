//go:build OMIT

package main

import (
	"fmt"
	"net/http"
)

// Exercice : faire tourner en local et ajouter quelques pages
// dont une avec un formulaire pour saisir un nom qui doit être
// affiché par la page hello

func hello(w http.ResponseWriter, r *http.Request) {
	nom := r.FormValue("nom") // Récupération du paramètre nom de la requête
	// astuce : le caractère ` permet de saisir une chaine sur plusieurs lignes
	// printf permet d'utiliser %s
	fmt.Fprintf(w, `<html>
	<body>
	Hello <b>%s</b>
	</body>
</html>`, nom)
}

func main() {

	http.HandleFunc("/hello", hello)

	fmt.Println("Listen :8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
