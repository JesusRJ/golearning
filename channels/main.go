package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type ticker struct {
	C chan struct{}
}

func main() {
	c := ticker{C: make(chan struct{}, 1)}

	go write(c.C)

	// fmt.Println("returned value: ", <-c.C)

	// t := time.NewTicker(time.Second * 1)

	ctx := context.Background()

	// select {
	// // case <-t.C:
	// // 	fmt.Println("Chegou aqui 1")
	// // case <-c.C:
	// // 	fmt.Println("Chegou aqui 2")
	// case <-ctx.Done():
	// 	fmt.Println("done")
	// default: // tipo um continuo do for
	// }

	for {
		select {
		// case <-t.C:
		// 	fmt.Println("Chegou aqui 1")
		case x, ok := <-c.C:
			if !ok {
				log.Fatal("deu merda no canal fechado")
			}
			fmt.Println("Chegou aqui 2 ", x)
		case <-ctx.Done():
			fmt.Println("done")
		}
	}
}

func write(c chan struct{}) {
	time.Sleep(1 * time.Second)
	c <- struct{}{}
	close(c)
}
