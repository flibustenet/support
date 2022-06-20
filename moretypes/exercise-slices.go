//go:build OMIT

package main

import (
	"fmt"
	"time"
)

var jours = []string{"Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi", "Dimanche"}

func main() {
	now := time.Now()
	now.Weekday
	for _, j := range jours {
		fmt.Println(j)
	}
}
