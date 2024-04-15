# Asserção de Tipo (Type Assertion)

Asserção de tipo (`Type assertion`) é uma operação que permite verificar se uma `variável de interface` coincide com determinado tipo e, caso verdade, extrair o valor concreto armazenado na variável.

> <details><summary style="color: green">Variável de interface</summary>
> 
> Variável de interface aponta pra uma interface, e não para uma implementação concreta. Sendo assim ele pode receber quaisquer tipos concretos que implementem a interface para a qual ela foi declarada. Desta forma dizemos que ela possuí um tipo dinâmico.
> </details>
>

Sintaxe:

```golang
x.(T)
```

Onde `x` (tipo dinâmico) é uma expressão de um tipo interface e `T` é um tipo (denominado tipo asserido/asserted). Uma asserção de tipo verifica se o tipo dinâmico de seu operando coincide com o tipo asserido.

Em uma asserção de tipo há duas possibilidades:

1. Se o tipo asserido `T` for **concreto**, a asserção verifica se o tipo dinâmico de `x` é idêntico a `T` e caso haja sucesso o resultado da asserção de tipo é o valor dinâmico de `x`, que claramente é do tipo `T`. Se a operação falhar será gerado um pânico (panic).

```golang
var w io.Writer
w = os.Stdout
f := w.(*os.File)     // sucesso: f == os.Stdout
c := w.(*byte.Buffer) // panic: a interface contêm *os.File e não *byte.Buffer
```

2. Se o tipo asserido `T` for um tipo **interface**, a asserção de tipo verifica se o tipo dinâmico de `x` satisfaz `T` e caso sucesso o valor dinâmico de `x` _não_ é extraído; o resultado continua sendo um valor interface contendo os mesmos componentes de tipo e valor, mas o resultado tem o tipo de interface `T`. A asserção de tipo para um tipo interface altera o tipo da expressão, deixando um conjunto diferente (geralmente maior) de métodos accessível, mas preserva os componentes de tipo e valor dinâmicos no valor da interface.

```golang
var w io.Writer         // io.Writer só possuí o método Writer
w = os.Stdout           // os.Stdout retorna um *os.File
rw := w.(io.ReadWriter) // sucesso: *os.File tem tanto o método Read quanto o método Write

w = new(ByteCounter)
rw  = (io.ReadWriter)   // panic: *ByteCounter não tem o método Read
```

> Independente do tipo asserido, a asserção de tipo irá falhar se o operando (`x`) for `nil`

Para evitar o lançamento de `panic`, a asserção de tipo permite o uso de [`atribuição de tupla`](../../atribuicoes/README.md#atribuição-de-tupla) para verificar o resultado da operação, retornando um segundo resultado booleano que indica sucesso.

```golang
var w io.Writer = os.Stdout

f, ok := w.(*bytes.Buffer)

if !ok {
  fmt.Println("w não é do tipo *bytes.Buffer")
}
```

Este comportamente é muito útil quando queremos testar se um tipo dinâmico é de algum tipo em particular.

# Referências

- [A Tour of Go - Type assertions](https://go.dev/tour/methods/15)
- [Go by example - Switch](https://gobyexample.com/switch)
