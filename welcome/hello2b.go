//go:build OMIT

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Exemple de lancement d'un serveur web

func main() {
	// déclaration de la fonction de manière anonyme
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello 世界")
	})

	// lancement du serveur dans une goroutine
	go func() {
		fmt.Println("Ecoute sur :8888")
		http.ListenAndServe(":8889", nil)
	}()

	// le programme continue dans la goroutine principale

	time.Sleep(time.Second) // attente lancement serveur
	// GET sur notre serveur
	resp, err := http.Get("http://localhost:8888")
	if err != nil {
		log.Fatal(err) // affichage d'une erreur éventuelle
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Reçu : %s\nErreur : %v\n", string(data), err)
}

// testez le cas où le serveur ne serait pas lancé sur le même port que le GET
// puis retirez le test d'erreur

// à voir par la suite : la gestion d'erreur particulière en Go
