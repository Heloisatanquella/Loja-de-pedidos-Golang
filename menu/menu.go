package menu

import "fmt"

func MenuPrincipal() {
	for {
		fmt.Print("\nMenu\n")
		fmt.Print("1 - Cadastrar novo cliente, produto ou pedido")
		fmt.Print("2 - Buscar clientes ou produtos")
		fmt.Print("3 - Sair\n")

		var opcao int
		_, err := fmt.Scanln(&opcao)

		if err != nil || opcao <= 0 && opcao >= 4 {
			fmt.Print("\nO valor inserido é inválido!\n")
		}
		if opcao > 0 && opcao < 4 {
			return
		}
		fmt.Print("\nO valor inserido é inválido!\n")
	}
}
