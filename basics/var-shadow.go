// go:build OMIT

package main

import (
	"fmt"
)

var GlobalVariable = "global"

func main() {
	loc := 8

	fmt.Println(loc, GlobalVariable)

	GlobalVariable := "ombre"

	fmt.Println("redef global", loc, GlobalVariable)

	{
		loc := 11
		fmt.Println("redef loc", loc, GlobalVariable)
	}
	fmt.Println("retour", loc, GlobalVariable)

}
