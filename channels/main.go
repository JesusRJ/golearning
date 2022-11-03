package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	write(c)

	fmt.Println("returned value: ", <-c)
}

func write(c chan<- int) {
	time.Sleep(2 * time.Second)
	c <- 10
}
