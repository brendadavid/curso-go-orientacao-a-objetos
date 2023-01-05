package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contaBrenda := ContaCorrente{titular: "Brenda", numeroAgencia: 1234, numeroConta: 56789, saldo: 150.55}
	contaBrenda2 := ContaCorrente{titular: "Brenda", numeroAgencia: 1234, numeroConta: 56789, saldo: 150.55}

	contaEduardo := ContaCorrente{"Eduardo", 6789, 45603, 200.0}
	contaEduardo2 := ContaCorrente{"Eduardo", 6789, 45603, 200.0}

	fmt.Println(contaBrenda == contaBrenda2)

	fmt.Println(contaEduardo == contaEduardo2)

	var contaCris *ContaCorrente
	contaCris = new(ContaCorrente)
	contaCris.titular = "Cris"
	contaCris.saldo = 55.0

	var contaCris2 *ContaCorrente
	contaCris2 = new(ContaCorrente)
	contaCris2.titular = "Cris"
	contaCris2.saldo = 55.0

	//false porque compara os endereços de memória
	fmt.Println(contaCris == contaCris2)

	//true porque está comparando o conteúdo dos objetos que são iguais
	fmt.Println(*contaCris == *contaCris2)

}
