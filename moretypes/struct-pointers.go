//go:build OMIT

package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	v := Point{1, 2}
	p := &v
	p.X = 13
	fmt.Println(v)
}
