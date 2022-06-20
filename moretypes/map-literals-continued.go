//go:build OMIT

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)

	t := map[string]struct {
		x int
		y int
	}{
		"a": {1, 2},
		"b": {2, 3},
	}
	fmt.Println(t)
}
