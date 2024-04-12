package main

import (
	"fmt"
	"reflect"
	"strings"
)

type (
	Person interface {
		Name() string
		Age() int
	}

	person struct {
		name    string
		age     int
		address []address
	}

	address struct {
		street string
		cep    string
		number int
	}
)

func (p person) Name() string {
	return p.name
}

func (p person) Age() int {
	return p.age
}

func main() {
	var p Person = person{
		name: "reginaldo jesus",
		age:  42,
		address: []address{
			{"address1", "", 0},
			{"address2", "", 0},
		},
	}

	fmt.Println(PrintStructAsCode(p))
}

func PrintStructAsCode(value any) string {
	v := reflect.ValueOf(value)
	t := v.Elem().Type()

	var result strings.Builder

	for x := 0; x < t.NumField(); x++ {
		fmt.Println(v.FieldByName(t.Field(x).Name))
		// result.WriteString
		// v := reflect.ValueOf(t.Field(x))

		// switch t.Field(x).Type.Kind() {
		// case reflect.Array:
		// case reflect.Slice:
		// case reflect.Struct:
		// 	fmt.Println("array")
		// }
	}

	return result.String()
}
