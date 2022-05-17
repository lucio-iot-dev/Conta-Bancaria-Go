package main

import (
	// "Conta-Bancaria-Go/clientes"
	"Conta-Bancaria-Go/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	  conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	 Sacar(valor float64) string
}

func main() {
	// clienteBruno := clientes.Titular{"Bruno", "123.111.123.12", "Desenvolvedor Go"}
  //  contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}
		
	// 	fmt.Println(contaDoBruno)

	// contaExemplo := contas.ContaCorrente{}
	// contaExemplo.Depositar(600)
	// fmt.Println(contaExemplo.ObterSaldo())

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)
	fmt.Println(contaDoDenis.ObterSaldo())
	
  contaDaLuisa := contas.ContaCorrente{}
	 contaDaLuisa.Depositar(500)
	 
	 PagarBoleto(&contaDaLuisa, 400)
	 fmt.Println(contaDaLuisa.ObterSaldo())

}
