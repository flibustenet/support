//go:build OMIT

package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	p := Point{1, 2}
	p.X = 4
	fmt.Println(p.X)
}
