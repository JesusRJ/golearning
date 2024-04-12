package main

import (
	"fmt"
	"sync"
)

// Este código testa o owner do channel.
// Se um channel somente leitura for criado em uma function, quem pode escrever no channel?

func main() {
	ch := make(chan any)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Write
	go func(c chan<- any) {
		defer wg.Done()
		defer close(c)
		for x := 0; x <= 5; x++ {
			c <- x
		}
	}(ch)

	// Read
	go func(c <-chan any) {
		defer wg.Done()
		for x := range c {
			fmt.Println("Value:", fmt.Sprintf("%v", x))
		}
	}(ch)

	wg.Wait()
}
