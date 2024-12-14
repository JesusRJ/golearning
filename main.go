package main

import "fmt"

func main() {
	fruits := map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"grape":  "purple",
	}

	for key := range fruits {
		fmt.Println(key)
	}
}
