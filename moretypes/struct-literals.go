//go:build OMIT

package main

import "fmt"

type Point struct {
	X, Y int
}

var (
	p0      Point          // initialisation à zéro
	p1      = Point{1, 2}  // has type Point
	p2      = Point{X: 1}  // Y:0 is implicit
	p3      = Point{}      // X:0 and Y:0
	pointer = &Point{1, 2} // has type *Point
	vp      *Point         // pointeur null
)

func main() {
	fmt.Println(p0, p1, pointer, p2, p3, vp)

	// struct anonyme
	p := struct {
		a int
		b int
	}{1, 2}
	fmt.Println(p)
}
