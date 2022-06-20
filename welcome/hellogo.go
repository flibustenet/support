//go:build OMIT

package main

import (
	"fmt"
	"time"
)

// Exemple du système de goroutines et communication par channel

func envoi(ch chan int) {
	tempo := time.Second * 2
	for i := 0; i < 5; i++ {
		fmt.Printf("Envoi %d dans %s\n", i, tempo)
		time.Sleep(tempo)
		ch <- i // envoi dans le chan
	}
	close(ch)
}

func recoit(ch chan int) {
	for res := range ch {
		fmt.Printf("    Reçoit %d\n", res)
	}
}

func main() {
	// création d'un canal de transmission : chan
	var ch = make(chan int)

	// envoi de messages dans le canal créé
	// dans une goroutine
	go envoi(ch)

	// reçoit dans la goroutine principale jusqu'à fermeture du canal
	recoit(ch)
}

// Changez la durée de la temporisation
// Déplacez la temporisation du côté réception

// Un canal permet à la fois de transmettre des données
// à la fois de synchroniser les goroutines

// La communication est uni-directionelle.
