package menu

import (
	"fmt"
	"loja-de-pedidos-golang/controllers"
	"loja-de-pedidos-golang/entities"
)

func MenuPrincipal() int {
	fmt.Println("=== Menu Principal ===")
	fmt.Println("1. Cadastro")
	fmt.Println("2. Busca")
	fmt.Println("3. Sair")
	fmt.Print("Escolha uma opção: ")

	var opcao int
	_, err := fmt.Scan(&opcao)
	if err != nil {
		fmt.Println("Erro ao ler a entrada, por favor digite um número.")
		opcao = 0
	}
	return opcao
}

func SubmenuCadastro() {
	for {
		fmt.Print("\n Submenu \n")
		fmt.Print("1 - Cadastrar novo cliente")
		fmt.Print("2 - Cadastrar novo produto")
		fmt.Print("3 - Cadastrar novo pedido a um novo cliente")
		fmt.Print("4 - Criar novos pedidos associados a clientes existentes")
		fmt.Print("5 - Voltar ao menu principal \n")

		var opcao int
		_, err := fmt.Scanln(&opcao)

		if err != nil || opcao <= 0 || opcao > 5 {
			fmt.Print("\nO valor inserido é inválido!\n")

			if opcao == 5 {
				break
			} else if opcao == 1 {
				controllers.CadastrarCliente(&[]entities.Cliente{})
			} else if opcao == 2 {
				controllers.Produto(&[]entities.Produto{})
			} else if opcao == 3 {
				controllers.PedidoCadastroCliente([]entities.Cliente{})
			} else if opcao == 4 {
				controllers.PedidoClienteExistente("email", []entities.Cliente{})
			} else {
				fmt.Print("\nO valor inserido é inválido!\n")
			}
		} else {
			fmt.Print("\nO valor inserido é inválido!\n")
		}

	}
}

func SubmenuBusca() {
	for {
		fmt.Print("\n Submenu \n")
		fmt.Print("1 - Buscar cliente")
		fmt.Print("2 - Buscar produto")
		fmt.Print("3 - Voltar ao menu principal \n")

		var opcao int
		_, err := fmt.Scanln(&opcao)

		if err != nil || opcao <= 0 || opcao > 5 {
			fmt.Print("\nO valor inserido é inválido!\n")

			if opcao == 3 {
				break
			} else if opcao == 1 {
				controllers.BuscarCliente("email", []entities.Cliente{})
			} else if opcao == 2 {
				controllers.BuscarProduto("nomeProduto", []entities.Produto{})
			} else {
				fmt.Print("\nO valor inserido é inválido!\n")
			}
		} else {
			fmt.Print("\nO valor inserido é inválido!\n")
		}
	}
}
