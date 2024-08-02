package wizard

import "fmt"

func Iniciar() {
	var opcao int

	fmt.Println("Bem-vindo ao sistema bancario IDM!")
	fmt.Printf("Escolha uma opção: \n1 - Acessar conta corrente\n2 - Criar conta corrente\n3 - Sair\n")

	fmt.Scanf("%d", &opcao)
	executarAcao(opcao)
}

func executarAcao(opcao int) {
	switch opcao {
	case 1:
		fmt.Println("Acessar conta corrente")
	case 2:
		fmt.Println("Criar conta corrente")
	case 3:
		fmt.Println("Saindo da aplicação...")
	default:
		fmt.Println("Opção inválida")
	}
}
