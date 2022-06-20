package main

import (
	"fmt"
	"log"
	"time"
)

/*
L'objectif est d'envoyer une série de mails sans attendre.
*/

func mails(ch chan string) {
	for msg := range ch {
		log.Printf("sendmail %s", msg)
		time.Sleep(time.Second) // le mail est lent à partir
	}
}
func main() {
	ch := make(chan string)
	go mails(ch)

	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("message %d", i)
		log.Printf("envoi %s", msg)
		ch <- msg
	}
	log.Println("suite...")
}
