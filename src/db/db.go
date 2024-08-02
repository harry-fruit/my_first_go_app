package db

import (
	conta_corrente "my-test/conta-corrente"
)

var contasCorrente []conta_corrente.TContaCorrente = []conta_corrente.TContaCorrente{
	{Titular: "João", Agencia: 1234, Conta: 123456, Saldo: 1000.00, Senha: "123456"},
	{Titular: "Maria", Agencia: 1234, Conta: 123457, Saldo: 2000.00, Senha: "123457"},
	{Titular: "José", Agencia: 1234, Conta: 123458, Saldo: 3000.00, Senha: "123458"},
	{Titular: "Pedro", Agencia: 1234, Conta: 123459, Saldo: 4000.00, Senha: "123459"},
	{Titular: "Ana", Agencia: 1234, Conta: 123460, Saldo: 5000.00, Senha: "123460"},
}

func GetContaCorrente(authEntrada conta_corrente.TAutenticacaoEntrada) (bool, conta_corrente.TContaCorrente) {
	existe := false

	for _, conta := range contasCorrente {
		if conta.Agencia == authEntrada.Agencia && conta.Conta == authEntrada.Conta && conta.Senha == authEntrada.Senha {
			existe = true
			return existe, conta
		}
	}

	return existe, conta_corrente.TContaCorrente{}
}
