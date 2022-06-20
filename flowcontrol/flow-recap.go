//go:build OMIT

package main

import "fmt"

func f() int {
	return 5
}

func main() {
	// initialisation et incrément dans la boucle
	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d\n", i)
		if i == 5 {
			fmt.Println(" i == 5")
			continue
		}
		if i == 8 {
			fmt.Println(" i == 8")
			break
		}
	}

	// initialisation et incrément en dehors de la boucle
	i := 0
	for i < 10 {
		fmt.Printf("i=%d\n", i)
		if i > 3 {
			goto label
		}
		i++
	}
label:

	// initialisation dans le if
	if s := f(); s == 5 {
		fmt.Println("f() == 5")
	}

	// switch sur les valeurs

	switch f() {
	case 0:
		fmt.Println("f == 0")
	case 1:
		fmt.Println("f == 1")
	case 5:
		fmt.Println("f == 5")
	default:
		fmt.Println("f <> 0, 1, 5")
	}

	// switch avec initialisation

	switch s := f(); s { // s est accessible au bloc
	case 0:
		fmt.Println("f == 0")
	default:
		fmt.Println("f ==", s)
	}

	// switch avec tests
	s := f()
	switch {
	case s == 0:
		fmt.Println("f == 0")
	case s == 5:
		fmt.Println("f == 5")
	}

}
