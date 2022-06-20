// go:build OMIT

package main

import "fmt"

func add(xs ...int) int {
	res := 0
	for _, x := range xs {
		fmt.Println("Add ", x)
		res += x
	}
	return res
}

func main() {
	fmt.Println(add(42, 13, 8, 3, 5))
}
