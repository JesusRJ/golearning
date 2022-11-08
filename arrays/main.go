package main

import "fmt"

func main() {
	// Declaração com tamanho definido
	var a [10]int
	fmt.Println("Length defined: ", a)

	// Declaração com tamanho inferido
	// Os três pontos (...) na declaração do Array diz ao Go para inferir o tamanho
	// do array a partir dos valores literais fornecidos. (Ref.: https://zetcode.com/golang/ellipses/)
	// Sem o três pontos seria retornado uma slice ao invés de um array.
	var b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("Length inferred: ", b)

	i := [10]int{1, 2}
	fmt.Println(i)
}
