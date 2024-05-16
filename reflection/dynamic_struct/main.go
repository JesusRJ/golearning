package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// Reference:
// https://github.com/knwoop/dynamic-struct/tree/main
// https://github.com/Chise1/dynamic-struct

func main() {
	structFields := []reflect.StructField{
		{Name: "Name", Type: reflect.TypeOf(""), Tag: reflect.StructTag(`json:"name"`)},
		{Name: "Age", Type: reflect.TypeOf(9), Tag: reflect.StructTag(`json:"age"`)},
		{Name: "Address", Type: reflect.TypeOf(""), Tag: reflect.StructTag(`json:"address"`)},
	}

	definitionType := reflect.StructOf(structFields)

	person := reflect.New(definitionType).Interface()

	data := `{"name": "Jesus", "age": 43, "address": "rua do além mar"}`

	if err := json.Unmarshal([]byte(data), person); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", person)
}
