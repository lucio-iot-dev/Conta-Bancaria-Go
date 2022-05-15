package main

import (
	// "Conta-Bancaria-Go/clientes"
	"Conta-Bancaria-Go/contas"
	"fmt"
)



func main() {
	// clienteBruno := clientes.Titular{"Bruno", "123.111.123.12", "Desenvolvedor Go"}
  //  contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}
		
	// 	fmt.Println(contaDoBruno)

	// contaExemplo := contas.ContaCorrente{}
	// contaExemplo.Depositar(600)
	// fmt.Println(contaExemplo.ObterSaldo())

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	contaDoDenis.Sacar(550)
	fmt.Println(contaDoDenis.ObterSaldo())
	


}
