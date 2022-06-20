//go:build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// renvoi une fonction
func more(m int) func(i int) bool {
	return func(i int) bool {
		return i > m
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	t := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		t[i] = rand.Intn(10)
	}

	fmt.Println("orig = ", t)
}
