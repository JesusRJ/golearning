package main

import (
	"fmt"
)

func span(numbers []int) (<-chan string, <-chan string, <-chan struct{}) {
	evenChan := make(chan string)
	oddChan := make(chan string)
	done := make(chan struct{})

	go func(numbers []int) {
		defer close(done)
		defer close(evenChan)
		defer close(oddChan)

		for _, value := range numbers {
			mod := value % 2
			if mod == 0 {
				oddChan <- fmt.Sprintf("Even: %d", value)
				continue
			}
			evenChan <- fmt.Sprintf("Odd: %d", value)
		}
	}(numbers)

	return evenChan, oddChan, done
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 17}
	evenChan, oddChan, done := span(numbers)

loop:
	for {
		select {
		case res := <-evenChan:
			fmt.Println(res)
		case err := <-oddChan:
			fmt.Println(err)
		case _, ok := <-done:
			if !ok {
				break loop
			}
		}
	}

	fmt.Println("Done")
}
