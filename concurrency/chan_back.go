//go:build OMIT

package main

import (
	"fmt"
	"time"
)

// back doit renvoyer un message quand elle a terminé
func back() chan bool {
	end := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(3 * time.Second)
			fmt.Println(i)
		}
		end <- true
	}()
	return end
}

func main() {
	end := back()
	for {
		select {
		case <-end: // attente de terminaison
			return
		default: // en attendant...
			fmt.Println("...")
			time.Sleep(time.Second)
		}
	}
}
