package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 50; i++ {
		fmt.Printf("%d ", i)
	}

	for {
		fmt.Println("\ninfinite loop")
		break
	}

	for i := 20; ; i-- {
		fmt.Printf("%d ", i)
		if i <= 0 {
			break
		}
	}

	fmt.Println("")

	for x := 1; ; x++ {
		mod := x % 5
		if mod == 0 {
			fmt.Printf("%d ", x)
			continue
		}
		if x > 100 {
			break
		}
	}
}
