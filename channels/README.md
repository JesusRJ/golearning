# Channels

esquema pub/sub

# CONCORRÊNCIA VS PARALELISMO (BOOK: THE PROGRAMMING GO)

- Canais sem buffer é obrigatório ter alguém ouvindo ele: sem isso podemos ter um dead-lock
- Canais bufferizados tem uma capacidade x de valores, que permite segurar um valor: tmb pode ocorrer um dead-lock se o buffer for totalmente preenchido, mas não for esvaziado. (ex. erros, signal)

# Fun facts

- Publicar pra channel não inicializado é uma operação blockante ad eterno
- Publicar pra channel fechado causa panic na aplicação
- Chamar close() em um canal fechado tmb causa panic
- \*\*Em um for-select eventualmente o go irá escolher um valor do canal caso eles estejam concorrendo/chegando ao mesmo tempo: por causa disso é usual termos selects dentro de selects para revalidarmos condições

# Tricks

```go
x, ok := <-c.C
if !ok {
  log.Fatal("deu merda no canal fechado")
}
```

- Se consultar um canal fechado ele retornar o valor zero do tipo do canal
