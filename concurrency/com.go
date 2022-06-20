//go:build OMIT

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Msg struct {
	Retour chan int
	Valeur int
}

func envoi(ctx context.Context, v chan Msg) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Envoi %d\n", i)
		res := make(chan int)
		j := i
		go func() {
			v <- Msg{res, j}
			fmt.Println("i=", j, "res=", <-res)
		}()
	}
}

func double(ctx context.Context, m chan Msg) {
	for {
		select {
		case <-ctx.Done():
			return
		case recep := <-m:
			go func() {
				res := recep.Valeur * 2
				rd := rand.Intn(5)
				fmt.Printf("wait %d pour %d\n", rd, recep.Valeur)
				time.Sleep(time.Duration(rd) * time.Second)
				recep.Retour <- res
			}()
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	reception := make(chan Msg)
	go double(ctx, reception)
	go envoi(ctx, reception)
	<-ctx.Done()
}
