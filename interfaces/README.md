# Interfaces

`Interfaces` fornecem flexibilidade de abstração para nossos programas.

Interfaces são `tipos abstratos` que definem um conjunto de funções que devem ser implementados para que um tipo seja considerado uma instância da interface. Elas fornecem uma forma de especificar que valores e ponteiros de um tipo em particular se comportem de determinada maneira.

Todos os métodos de um tipo de interface são considerados a interface.

> Para um tipo satisfazer uma interface ele deve implementar todos os métodos requeridos por aquela interface.

Uma `interface` é duas coisas: um conjunto de métodos (method set) e um tipo (type).

A maior vantagem de usar interfaces é que você pode passar uma variável que implementa uma interface em particular para qualquer função que espera um parâmetro daquela interface específica, o que permite escrever funções mais flexíveis e adaptáveis que não dependam de uma implementação única.

# Referências

- [Methods, Interfaces and Embedded Types in Go](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html)
- [Livro: Mastering Go: Create Golang production applications using network libraries, concurrency, machine learning, and advanced data structures](https://www.amazon.com.br/Mastering-production-applications-concurrency-structures-ebook/dp/B07WC24RTQ/ref=sr_1_2?crid=1PHHVUHJ5XC8&keywords=mastering+go&qid=1669081526&qu=eyJxc2MiOiIxLjg5IiwicXNhIjoiMS4wMCIsInFzcCI6IjAuMDAifQ%3D%3D&sprefix=mastering+go%2Caps%2C215&sr=8-2&ufe=app_do%3Aamzn1.fos.4bddec23-2dcf-4403-8597-e1a02442043d)
