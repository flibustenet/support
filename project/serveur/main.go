package main

import (
	"fmt"
	"log"
	"net/http"

	"serveur.flib.fr/hdl"
)

// Exemple de lancement d'un serveur web
func main() {
	http.HandleFunc("/", hdl.Index)
	http.HandleFunc("/hello", hdl.Hello)

	fmt.Println("Listen :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
