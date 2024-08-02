package conta_corrente

import (
	"banco_app/pkg/conta"
	"fmt"
)

type ContaCorrente struct {
	conta.Conta
	saldo float64
}

func (c *ContaCorrente) Depositar(valorDeposito float64) {

	if valorDeposito <= 0 {
		fmt.Println("Valor de depósito inválido - O valor de depósito precisa ser maior que 0.")
		return
	}

	c.saldo += valorDeposito
	fmt.Println("Depósito efetuado com sucesso.")

}

func (c *ContaCorrente) Sacar(valorSaque float64) {

	if valorSaque <= 0 {
		fmt.Println("Valor de saque inválido - O valor de saque precisa ser maior que 0.")
		return
	}

	if valorSaque > c.saldo {
		fmt.Println("Saldo insuficiente.")
		return
	}

	c.saldo -= valorSaque
	fmt.Println("Saque efetuado com sucesso.")

}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) {

	if valorTransferencia <= 0 {
		fmt.Println("Valor de transferência inválido - O valor de transferência precisa ser maior que 0.")
		return
	}

	if valorTransferencia > c.saldo {
		fmt.Println("Saldo insuficiente.")
		return
	}

	c.saldo -= valorTransferencia
	contaDestino.saldo += valorTransferencia
	fmt.Println("Transferência efetuada com sucesso.")

}
