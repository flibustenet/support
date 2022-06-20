//go:build OMIT

package main

import (
	"fmt"
	"net/http"
)

// Exemple de lancement d'un serveur web
// La lib standard inclue les functions nécessaires

func hello(w http.ResponseWriter, r *http.Request) {
	nom := r.FormValue("nom") // Récupération du paramètre nom de la requête
	fmt.Fprintf(w, "<html><body>Hello <b>"+nom+"</b></body></html>")
}

func main() {

	http.HandleFunc("/hello", hello)

	fmt.Println("Listen :8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
