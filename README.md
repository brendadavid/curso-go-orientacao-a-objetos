# Anotações sobre o código

Uma `struct` recebe sempre o prefixo `type` e o sufixo `struct`, é criado fora da main e é parecido com um construtor de Java

Se todos as variáveis forem utilizadas, não precisa informar o nome da variável na criação do "objeto", caso só algumas sejam utilizadas, as outras serão inicializadas com o valor default para cada tipo.

A `struct` pode ser inicializada com a palavra `new(nomeStruct)`, mas o nome da struct precisa receber um **"*"** antes, para o Go saber para onde o ponteiro deve apontar.