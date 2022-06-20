//go:build OMIT

package main

import (
	"fmt"
)

func Div(x, y float64) float64 {
	return x / y
}

func main() {
	fmt.Println(Div(4, 2))
	fmt.Println(Div(8, 2))
	fmt.Println(Div(8, 0))
}
