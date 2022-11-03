package main

import "fmt"

func main() {
	var v uint8

	for {
		v++
		fmt.Println(v)
		if v > 255 {
			break
		}
	}

}
