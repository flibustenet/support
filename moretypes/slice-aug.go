//go:build OMIT

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	a = append(a, 5)
	b := a[:]

	fmt.Println("a=", a, "b=", b)
	fmt.Println("cap(a)=", cap(a), "cap(b)=", cap(b))

	a[0] = 9 // même array sous-jacent
	fmt.Println("a=", a, "b=", b)
	fmt.Println("cap(a)=", cap(a), "cap(b)=", cap(b))

	a = append(a, 5, 6, 7, 8)
	a[0] = 99 // peut-être plus le même array
	fmt.Println("a=", a, "b=", b)
	fmt.Println("cap(a)=", cap(a), "cap(b)=", cap(b))
}
