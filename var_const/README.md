# Variáveis e Constantes

Referências:

- [Como usar variáveis e constantes em Go](https://www.digitalocean.com/community/tutorials/how-to-use-variables-and-constants-in-go-pt)

# Nomeando variáveis: regras e estilo

A nomeação de variáveis é bastante flexível, mas existem algumas regras a serem lembradas:

- Os nomes das variáveis devem consistir em apenas uma palavra (ou seja, sem nenhum espaço).
- Os nomes das variáveis devem ser constituídos apenas de letras, números e sublinhados (\_).
- Os nomes das variáveis não podem começar com um número.
- Os nomes das variáveis são `cases sensitives` - evite o uso de nomes de variáveis semelhantes dentro de um programa.
- Os nomes de variáveis iniciadas com letra em caixa alta são variáveis exportadas e acessíveis fora do pacote em que foram declaradas. Se uma variável iniciar com uma letra em caixa baixa, então ela está disponível apenas dentro do pacote em que foi declarada.
- Quanto menor o escopo em que a variável existe, menor o nome da variável:

```golang
names := []string{"Mary", "John", "Bob", "Anna"}
for i, n := range names {
	fmt.Printf("index: %d = %q\n", i, n)
}
```

`names` tem um escopo mais amplo e possuí um nome mais significativo. `i` e `n` são usados em um escopo mais restrito, dentro do loop somente, e recebem um nome reduzido.

- Use `MixedCaps` ou `mixedCaps` ao invés de sublinhados para nomes com várias palavras.
- Acrônimos devem ser em letras maiúsculas. (ex.: `serveHTTP`).

# Valor Zero

Em Go todos os tipos têm um valor zero, não podemos ter valores `undefined` [não definidos] como em algumas outras linguagens. Por exemplo, um **boolean** em algumas linguagens poderia ser `undefined`, `true`, ou `false`, o que permite três estados para a variável. Em Go, não podemos ter mais de dois estados de um valor booleano.
