package main

import "fmt"

var Buttons byte

func main() {
	Buttons += 255 << 0
	fmt.Printf("%08b\n", Buttons)
}
