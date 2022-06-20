//go:build OMIT

package main

import "fmt"

func main() {
	i := 0
	defer fmt.Println("defer i=", i)
	i++
	fmt.Println("i=", i)
}
