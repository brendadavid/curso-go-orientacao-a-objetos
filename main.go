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
	contaEduardo := ContaCorrente{"Eduardo", 6789, 45603, 200.0}

	fmt.Println(contaBrenda)
	fmt.Println(contaEduardo)

	var contaCris *ContaCorrente
	contaCris = new(ContaCorrente)
	contaCris.titular = "Cris"
	contaCris.saldo = 55.0

	fmt.Println(*contaCris)

}
