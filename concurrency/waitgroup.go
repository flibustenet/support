//go:build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	max := 8
	wg.Add(max)
	go func() {
		for i := 0; i < max; i++ {
			time.Sleep(time.Second)
			fmt.Print(".")
			wg.Done()
		}
	}()

	wg.Wait()
}
