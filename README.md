# Anotações sobre o código

Uma `struct` recebe sempre o prefixo `type` e o sufixo `struct`, é criado fora da main e é parecido com um construtor de Java

Se todos as variáveis forem utilizadas, não precisa informar o nome da variável na criação do "objeto", caso só algumas sejam utilizadas, as outras serão inicializadas com o valor default para cada tipo.

A `struct` pode ser inicializada com a palavra `new(nomeStruct)`, mas o nome da struct precisa receber um `"*"` antes, para o Go saber para onde o ponteiro deve apontar. Para printar o conteúdo da variável apenas, o `"*"` também deve ser inserido o antes da variável.

Em uma comparação, a `struct` iniciada sem a palavra *new* pode ser comparada diretamente, o Go entende que os conteúdos é que estão sendo comparados.

### A saída será true em ambos pois seus conteúdos são iguais

    fmt.Println(contaBrenda == contaBrenda2)

    fmt.Println(contaEduardo == contaEduardo2)

Quando são inicializados com a palavra chave *new*, a menos que seja colocado um "*" antes das variáveis, o Go irá comparar os endereços de memódia e a saída será false.

	//false porque compara os endereços de memória
	fmt.Println(contaCris == contaCris2)

	//true porque está comparando o conteúdo dos objetos que são iguais
	fmt.Println(*contaCris == *contaCris2)

Deve ser adicionado um predicado + *nomeStruct entre parênteses antes do nome da função para que o Go saiba qual objeto está chamando a função e para onde o ponteiro deve apontar. Ex: `func (c *ContaCorrente) Sacar(valorDoSaque float64)`, para a contaSilvia os atributos dentro da função serão acessados como c.saldo, c.titular..

É possível que uma função receba um número indeterminado de parâmetros. Funções deste tipo são conhecidas em Go como variadic functions, ou função variádica. Para criar uma variadic function, devemos preceder o tipo do argumento com reticências, ex: 
	
	func Somando(numeros ...int) int

	func main() {
		fmt.Println(Somando(1))
		fmt.Println(Somando(1,1))
	}

