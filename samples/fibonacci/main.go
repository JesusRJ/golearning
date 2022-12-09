package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please: Inform the required argument number!")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Required a valid number as argument!")
	}

	fmt.Printf("The Fibonacci Sequence for \"%d\" is: %d\n", n, fibonacci(n))
}

func fibonacci(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
