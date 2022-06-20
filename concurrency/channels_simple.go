//go:build OMIT

package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
	}()
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y)
}
