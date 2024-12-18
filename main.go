package main

import (
	"fmt"
	"loja-de-pedidos-golang/menu"
)

func main() {
	for {
		opcao := menu.MenuPrincipal()
		switch opcao {
		case 1:
			menu.SubmenuCadastro()
		case 2:
			menu.SubmenuBusca()
		case 3:
			fmt.Println("Encerrando programa...")
			return
		default:
			fmt.Println("Opção inválida, tente novamente.")
		}

	}
}
