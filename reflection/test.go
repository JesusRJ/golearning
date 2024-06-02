package main

import (
	"fmt"
	"reflect"
)

type MyFloat float64

type MyStruct struct {
	Value string
}

func main() {
	// var x MyFloat = 35.5
	// myStruct := MyStruct{"Hello"}

	var myInterface interface{} = "OI"

	v := reflect.ValueOf(myInterface)

	fmt.Println(v.Type())
	fmt.Println(v.Kind())
	if v.Kind() == reflect.Ptr {
		fmt.Println(".Elem() ->", v.Elem())
		// fmt.Println(rValue.Elem())
	}
	fmt.Println(v.Interface())
}
