//go:build OMIT

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for i := 0; i < 3; i++ {
		j := i
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Println("-", j)
					time.Sleep(time.Second)
				}
			}
		}(ctx)
	}
	for i := 0; i < 10; i++ {
		fmt.Println("...", i)
		time.Sleep(time.Second)
	}
}
