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
	}

	definitionType := reflect.StructOf(structFields)

	// With set field
	value := reflect.New(definitionType)
	value.Elem().FieldByName("Name").SetString("JESUS")
	value.Elem().FieldByName("Age").SetInt(43)
	fmt.Printf("%+v\n", value)

	// With JSON Unmarshal
	person := reflect.New(definitionType).Interface()
	data := `{"name": "Miguel", "age": 9, "address": "rua do além mar"}`
	if err := json.Unmarshal([]byte(data), person); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", person)

}
