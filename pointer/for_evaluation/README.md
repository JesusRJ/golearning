# Pointer Evaluation em For (Avaliação de Ponteiros)

Quando se usa um ponteiro como variáveil de controle em um loop `for`, a avaliação do ponteiro ocorre em momentos distintos.
Isto causa que todas as referências internas no loop apontem para o mesmo lugar.

Para resolver isso é necessário realizar o reassing para uma variável interna ao loop.
