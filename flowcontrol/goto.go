//go:build OMIT

package main

import "fmt"

func main() {

	for i := 0; ; i++ {
		fmt.Println(i)
		if i >= 5 {
			goto lbl // remplacez par break
		}
	}
	fmt.Println("Fin de bloc")
lbl:
	fmt.Println("Après le label")
}
