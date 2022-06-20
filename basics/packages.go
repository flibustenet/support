// go:build OMIT

package main

import (
	"fmt"
	"strings"
)

func Majuscule(s string) string {
	return strings.ToUpper(s)
}

func main() {
	fmt.Println(Majuscule("hello"))
}
