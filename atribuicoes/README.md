# Atribuições

Uma instrução de atribuição atualiza o valor mantido por uma variável.

Em Go usamos o sinal `=` para atribuir um valor á uma variável:

```golang
x = 10                // variável nomeada
y = x * 5
a[0] = 5              // elemento de um array, slice ou map
*p = 100              // variável indireta
person.Name = "Jesus" // campo de uma struct
```

De forma simples, do lado esquerdo temos uma variável e do lado direito temos o valor ou expressão que retorna um valor.

## Atribuição implícita e explicíta

Há muitos outros lugares em um programa onde uma atribuição ocorre **implicitamente**:

- Uma chamada de função atribui implicitamente os valores de argumentos as variáveis de parâmetos correspondentes;
- A instrução `return` atribui implicitamente os operandos de retorno às variáveis de resultados correspondentes;
- Uma expressão literal para um tipo composto atribui implicitamente cada elemento como se fosse escrito separadamente:

```golang
medals := []string{"gold", "silver", "bronze"}
```

pode ser visto como:

```golang
medals[0] = "gold"
medals[1] = "silver"
medals[2] = "bronze"
```
- Mapas e canais, embora não sejam variáveis comuns, também estão sujeitos a atribuições implicitas semelhantes;

> :mega: Independente de ser explícita ou implícita, uma atribuição é sempre permitida se o lado esquerdo e direito forem do mesmo tipo: a atribuição só é permitida se o valor puder ser atribuído ao tipo da variável.

## Atribuição de Tupla

A `atribuição de tupla` permite que diversas variáveis recebam valores de uma única vez.

> :mega: Todas as expressões do lado direito são avaliadas antes que qualquer variável do lado esquerdo seja atualizada.

É muito útil quando quando algumas das variáveis aparecem em ambos os lados da atribuição, como por exemplo ao fazer um swap de variáveis (troca de valores):

```golang
x := 10
y := 5

x, y := y, x
```

Também pode deixar uma sequência de atribuições triviais mais compacta:

```golang
x, y, z = 1, 2, 3
```

> :exclamation: Por questões de estilo, você deve evitar a forma de tuplas se as expressões forem complexas, pois uma sequência de instruções separadas é mais fácil de ler.

Algumas expressões (funções, [type assertion](../type_assertion/README.md), [channels](../channels/README.md)(asserção de tipo)) usam a forma de tupla para retornar mais de um valor. Quando uma chamada desse tipo é usada em uma instrução de atribuição, o lado esquerdo deve conter tantas variáveis quanto forem os resultados da expressão:

```golang
f, err := os.Open("teste.txt")
```

Constantemente funções utilizam este recurso para retornar um resultado adicional para indicar a ocorrência de um erro na função, ou um `bool` (normalmente chamado `ok`) para indicar o resultado com sucesso de uma expressão.