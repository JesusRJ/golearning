# Golang - Guia de Aprendizado

<div>
<img src="./assets/desk.svg" alt="drawing" style="width:20%;"/>Este projeto contém notas de estudo da linguagem Golang.
</div>

# Dicas

## `fmt.Printf`

- O verbo `%T` imprime o `data type` da variável e evita a necessidade de usar o pacote `reflect`.

```golang
var d bool
fmt.Printf("var a %T = %+v\n", a, a)
```

# Referências

- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) - Tips and conventions the community follows
