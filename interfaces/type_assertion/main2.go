package main

import (
	"fmt"
	"log"
)

type Pessoa interface {
	print()
}

type PessoaFisica struct {
	Name string
	Age  uint
}

type PessoaJuridica struct {
	RazaoSocial string
	CNPJ        string
}

func (p PessoaFisica) print() {
	fmt.Println("Person", p.Name, "has", p.Age)
}

func main() {
	var pf Pessoa
	pf = PessoaFisica{
		Name: "Reginaldo Jesus",
		Age:  41,
	}

	p, ok := pf.(PessoaFisica)
	if !ok {
		log.Fatal("Invalid person")
	}
	p.print()

	var n interface{} = nil
	a := n.(Pessoa)
	a.print()

}
