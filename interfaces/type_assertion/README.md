# Asserção de Tipo (Type Assertion)

Asserção de tipo (`Type assertion`) é uma operação que permite verificar se uma `variável de interface` coincide com determinado tipo.

> <details><summary style="color: green">Variável de interface</summary> aponta pra uma interface, e não para uma implementação concreta. Sendo assim ele pode receber quaisquer tipos concretos que implementem a interface para a qual ela foi declarada. Desta forma dizemos que ela possui um tipo dinâmico.
> 
> <p>
> 
> #### We can hide anything, even code!
> 
> ```ruby
>   puts "Hello World"
> ```

</p>
</details>

```golang
x.(T)
```

Só pode ser usado com `interfaces`.

# Referências

- [A Tour of Go - Type assertions](https://go.dev/tour/methods/15)
- [Go by example - Switch](https://gobyexample.com/switch)
