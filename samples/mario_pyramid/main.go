package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func validateInput(v string) (int, bool) {
	r, err := strconv.Atoi(v)
	if err != nil {
		return 0, false
	}

	if r < 0 || r > 8 {
		return 0, false
	}

	return r, true
}

func printPyramid1(height int) {
	for x := 0; x < height; x++ {
		l1 := strings.Repeat(" ", height-x) + strings.Repeat("#", x+1)
		l2 := strings.Repeat("#", x+1) + strings.Repeat(" ", height-x)
		fmt.Println(l1 + "  " + l2)
	}
}

func main() {
	var height int

	for {
		fmt.Print("Height: ")

		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal(err)
		}

		var ok bool
		height, ok = validateInput(input)
		if ok {
			break
		}
	}

	printPyramid1(height)
}
