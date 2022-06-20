//go:build OMIT

package main

import (
	"fmt"
)

func imprevisible(i int) string {
	return []string{"zero", "un", "deux"}[i]
}

func appel(i int) (ret string) {
	defer func() {
		if e := recover(); e != nil {
			ret = fmt.Sprintf("Erreur : %v", e)
		}
	}()
	ret = imprevisible(i)
	return
}

func main() {
	res := appel(1)
	fmt.Println("res=", res)
}
