//go:build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Exercice : trier le slice avec les nombres triés, pairs en premier
// 3 2 1 5 4 -> 2 4 1 3 5

func main() {
	rand.Seed(time.Now().UnixNano())
	a := []int{}
	for i := 0; i < 20; i++ {
		a = append(a, rand.Intn(10))
	}
	fmt.Println("init a=", a)

	fmt.Println("a=", a)
}
