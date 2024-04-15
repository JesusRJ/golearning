package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second * 2)
			fmt.Println("Ok")
		}()
	}

	wg.Wait()

	fmt.Println("Finish")
}
