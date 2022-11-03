package main

import "fmt"

/*
Funções variádicas são funções que recebem qualquer número de parâmetros.
Variadic em Go são denotados pela notação de três pontos "...".

Verifique em ../arrays para saber como é possível utilizar a notação de três pontos
para determinar o tamanho de um array literal.

Ref.: https://yourbasic.org/golang/three-dots-ellipsis/
*/

func main() {
	nomes("Reginaldo", "José", "Jesus")
	fmt.Println("Soma: ", sum(5, 6, 9, 10))

	avgByGroup("G1", 15, 16, 11, 12, 17)

	// É possível passar os valores de uma slice diretamente para uma função variádica
	// se você desempacotar ela com "..."
	ages := []int{23, 25, 30, 25, 20, 21, 22, 26}
	avgByGroup("G2", ages...)
}

// Função variádica que recebe várias strings e itera sobre elas.
func nomes(s ...string) {
	for _, x := range s {
		fmt.Println(x)
	}
}

// Função variádica que soma vários números
func sum(n ...int) int {
	var r int
	for _, x := range n {
		r = r + x
	}
	return r
}

// Somente o último parâmetro de uma função pode ser variável
func avgByGroup(group string, ages ...int) {
	var t int
	for _, x := range ages {
		t = t + x
	}
	fmt.Println(group, ":", t/len(ages))
}
