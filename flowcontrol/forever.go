//go:build OMIT

package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	for i := 0; ; i++ {
		if i < 3 {
			continue
		}
		fmt.Println("i=", i)
		if i > 5 {
			break
		}
	}

	go func() {
		time.Sleep(3 * time.Second)
		log.Fatalf("stop")
	}()

	for {
		time.Sleep(time.Second)
		fmt.Print(".")
	}
}
