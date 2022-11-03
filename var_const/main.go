package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("\nVariables ====================")
	variables()
	mutualAtribuitions()

	fmt.Println("\nConstants ====================")
	constants()
}

// Variables
func variables() {
	// Valores zero: todo tipo tem um valor zero atribuído
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a %T = %+v\n", a, a)
	fmt.Printf("var b %T = %q\n", b, b)
	fmt.Printf("var c %T = %+v\n", c, c)
	fmt.Printf("var d %T = %+v\n\n", d, d)
}

func mutualAtribuitions() {
	j, k, l := "shark", 2.05, 15
	fmt.Printf("%s (is %T)\t", j, j)
	fmt.Printf("%f (is %T)\t", k, k)
	fmt.Printf("%d (is %T)\n", l, l)
}

// Constants
const text = "Jesus"

func constants() {
	fmt.Printf("%s (is %T)\n", text, text)

	const n = 500000000
	fmt.Printf("%d (is %T)\n", n, n)

	const d = 3e20 / n
	fmt.Printf("%f (is %T)\n", d, d)

	fmt.Printf("(is %T)\n", d)
	fmt.Println(math.Sin(n))
}
