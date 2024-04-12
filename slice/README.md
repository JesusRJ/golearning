# Slices

> Veja: https://go.dev/blog/slices-intro

Em Go, uma `Slice` é uma estrutura de dados dinâmica construída sobre um array.
Uma `slice` é definida por um tipo que inclui um tipo de elemento e um conjunto de limites. Por exemplo, `[]int` é uma fatia de inteiros e `[]string` é uma fatia de strings.

Quando você define uma variável de `slice` em Go, como var `mySlice []int`, Go cria uma estrutura de dados que contém três informações:

1. Um ponteiro para um array subjacente que contém os elementos da slice
2. Um comprimento que representa o número de elementos na slice
3. Uma capacidade que representa o número máximo de elementos que o array subjacente pode conter

Por exemplo, se você criar uma fatia como esta:

```golang
mySlice := []int{ 2, 3,4}
```

Go cria um array subjacente como este:

```
+---------+---+
| 2| 3| 4|
+---------+- -- +
len=3, cap=< ai=15>4
```

O cabeçalho da `slice` contém um ponteiro para o início da matriz (neste caso, o segundo elemento), bem como o comprimento e a capacidade da `slice`.

Veja o que há em um cabeçalho de uma `slice` verificando o [tipo reflect.SliceHeader](https://pkg.go.dev/reflect#SliceHeader):

```golang
type SliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
}
```

## Slices passados como valor Vs Pointeiro

Se passarmos a fatia para uma função como parâmetro:

```golang
func myFunc(s []int) {
  // Faça algo com a fatia(mySlice)myFunc
}
```

Go cria uma cópia do cabeçalho da fatia e a passa para a função. A cópia tem o mesmo comprimento, capacidade e ponteiro que a fatia original.

```
+---------+
| 2 | 3 | 4 |
+---------+
  len=3, cap=4

Copy of slice header:
+---------+
| 2 | 3 | 4 |
+---------+
  len=3, cap=4
```

No entanto, alterações na cópia, não alteram a `slice` passada como referência, mas alterações no valor do array, afetam o array subjacente:

```golang
package main

import "fmt"

func myFunc(s []int) {
    s[0] = 10
    s = append(s, 6)
}

func main() {
    mySlice := []int{1, 2, 3, 4, 5}
    myFunc(mySlice)
    fmt.Println(mySlice)
}

Output : [10 2 3 4 5]
```

A operação de modificar o primeiro elemento da slice na função `myFunc(s []int)`, altera o array subjacente, mas a operação `s = append(s, 6)` não tem efeito sobre a slice, uma vez que estamos atuando em uma cópia do cabeçalho dela.

Se quisermos que myFunc modifique/anexe a fatia existente, precisamos usar o ponteiro:

```golang
package main

import "fmt"

func myFunc(s *[]int) {
    (*s)[0] = 10
    *s = append(*s, 6)
}

func main() {
    mySlice := []int{1, 2, 3, 4, 5}
    myFunc(&mySlice)
    fmt.Println(mySlice)
}

Output : [10 2 3 4 5 6]
```

# Referências

- https://go.dev/blog/slices-intro
- https://medium.com/@ansujain/understanding-slices-in-go-pass-by-value-and-pass-by-pointer-9a52830c741e
