//go:build OMIT

package main

import (
	"fmt"
)

type Entier int

func (e Entier) Double() Entier {
	return e * 2
}
func (e Entier) String() string {
	return fmt.Sprintf("valeur=%d", e)
}

func Add(x, y int) int {
	return x + y
}

func main() {
	e := Entier(5)
	e += 1
	fmt.Println(e, e.Double())

	fmt.Println(Add(int(e), 3))
}
