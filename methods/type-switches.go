//go:build OMIT

package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case V:
		fmt.Printf("V type = %T V.x=%d", v, v.x)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type V struct {
	x int
}

func main() {
	do(21)
	do("hello")
	do(true)
	do(V{5})
}
