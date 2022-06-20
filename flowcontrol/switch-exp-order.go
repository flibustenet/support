//go:build OMIT

package main

import (
	"fmt"
)

func main() {
	// ordre d'évaluation des listes d'expressions
	switch Foo(2) {
	case Foo(1), Foo(2), Foo(3):
		fmt.Println("First case")
		fallthrough
	case Foo(4):
		fmt.Println("Second case")
	}
}
func Foo(n int) int {
	fmt.Println(n)
	return n
}
