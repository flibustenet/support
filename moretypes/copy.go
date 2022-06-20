//go:build OMIT

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 3, 5)
	fmt.Printf("Copie %d éléments\n", copy(b, a))
	fmt.Println("a=", a)
	fmt.Println("b=", b)
}
