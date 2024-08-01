package conta_corrente

import (
	"fmt"
	"my-test/db"
)

type TAutenticacaoSaida struct {
	autenticado bool
	mensagem    string
}

type TAutenticacaoEntrada struct {
	Conta   int
	Agencia int
	Senha   string
}

func autenticar() TAutenticacaoSaida {
	conta := solicitarConta()
	agencia := solicitarAgencia()
	senha := solicitarSenha()

	db.GetContaCorrente()
}

func solicitarAgencia() int {
	var agencia int
	fmt.Print("Agência: ")
	fmt.Scan("%d", &agencia)
	return agencia
}

func solicitarConta() int {
	var conta int
	fmt.Print("Conta: ")
	fmt.Scan("%d", &conta)
	return conta
}

func solicitarSenha() string {
	var senha string
	fmt.Print("Senha: ")
	fmt.Scan("%s", &senha)
	return senha
}
