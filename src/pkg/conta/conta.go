package conta

import (
	"banco_app/pkg/individuo"
	"fmt"
)

type Conta struct {
	Agencia, Conta int
	Individuo      individuo.Individuo
	senha          string
}

func (c *Conta) SetSenha(senha string) {
	c.senha = senha
	fmt.Println("Senha cadastrada com sucesso!")
}
