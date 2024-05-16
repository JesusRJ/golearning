package main

import (
	"fmt"
	"reflect"
)

// Reference: https://github.com/knwoop/dynamic-struct/tree/main

func main() {
	structFields := []reflect.StructField{
		{Name: "Name", Type: reflect.TypeOf(3), Tag: reflect.StructTag("json:\"name\"")},
		{Name: "Age", Type: reflect.TypeOf(9), Tag: reflect.StructTag("json:\"age\"")},
		{Name: "Address", Type: reflect.TypeOf(3), Tag: reflect.StructTag("json:\"address\"")},
	}

	definitionType := reflect.StructOf(structFields)

	person := reflect.New(definitionType).Interface()

	fmt.Printf("%+v", person)
}
