//go:build OMIT

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := []int{}
	for i := 0; i < 10; i++ {
		a = append(a, rand.Intn(10))
	}
	fmt.Println("init a=", a)

	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	fmt.Println("rand a=", a)

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

	fmt.Println("sort a=", a)
}
