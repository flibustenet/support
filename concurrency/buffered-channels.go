//go:build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5) // changer la taille du tampon
	// ch := make(chan int) // sans tampon
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("Reçoit ", <-ch)
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Printf("Envoi %d\n", i)
		ch <- i
	}
	time.Sleep(time.Second * 6)
}
