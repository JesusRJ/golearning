package main

import (
	"fmt"
	"reflect"
)

type Pessoa struct {
	Name string
	CPF  string
}

func main() {
	cliente := Pessoa{
		Name: "Mario José Cortês",
		CPF:  "123.432.456-01",
	}

	addPrefixType(&cliente)

	fmt.Printf("%+v", cliente)
}

func addPrefixType(p *Pessoa) {
	val := reflect.ValueOf(p)

	nameField := reflect.Indirect(val).FieldByName("Name")
	nameField.SetString("TESTE")
}
