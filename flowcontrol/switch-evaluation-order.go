//go:build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Quand arrive dimanche ?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Aujourd'hui")
	case today + 1:
		fmt.Println("Demain.")
	case today + 2:
		fmt.Println("Dans 2 jours")
	case today + 3, today + 4:
		fmt.Println("Dans 3 ou 4 jours")
	default:
		fmt.Println("Dans longtemps")
	}
}
