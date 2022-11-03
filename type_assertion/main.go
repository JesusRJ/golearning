package main

import "fmt"

func main() {
	typeAssertion()
	typeSwitch()
}

// Execute type assertion directly
func typeAssertion() {
	var i interface{} = 500
	t, ok := i.(int)
	if ok {
		fmt.Printf("The data %d is from %T type\n", i, t)
	}
}

// Execute type assertions using `type switch`
func typeSwitch() {
	var i interface{} = 500

	switch t := i.(type) {
	case bool:
		fmt.Println("Is a bool")
	case int:
		fmt.Println("Is a int")
	default:
		fmt.Printf("Unknow type %T\n", t)
	}

	whatIs := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Is a bool")
		case int:
			fmt.Println("Is a int")
		default:
			fmt.Printf("Unknow type %T\n", t)
		}
	}

	whatIs(true)
	whatIs(500)
	whatIs(.5)
}
