package main

import (
	"fmt"
	"reflect"
)

type MyFloat float64

func main() {

	var x MyFloat = 35.5

	p := reflect.ValueOf(&x)

	fmt.Println(p.Type())
	fmt.Println(p.Kind())
	fmt.Println(p.Elem())

	fmt.Println(p.Interface())
}
