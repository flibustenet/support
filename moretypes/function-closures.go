//go:build OMIT

package main

import "fmt"

func adder(sum int) func(int) int {
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos := adder(5)
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
	}
}
