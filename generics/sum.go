//go:build OMIT

package main

import (
	"fmt"
)

type S string

type Number interface {
	int64 | float64 | int
}

func sum[V Number | ~string](m []V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
func main() {
	fmt.Println(sum([]int{1, 2, 3}))
	fmt.Println(sum([]string{"a", "b", "c"}))
	fmt.Println(sum([]S{"a", "b", "c"}))
	fmt.Println(sum([]float64{0.5, 1, 2}))
}
