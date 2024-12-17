package controllers

import (
	"fmt"
	"loja-de-pedidos-golang/entities"
)

func Pedidos(cliente entities.Cliente, produtosEmEstoque *[]entities.Produto, listaPedidos *[]entities.Pedido) *entities.Pedido {
	var itensPedido []entities.ItemPedido
	var totalDoPedido float64
	clienteEmail := cliente.Email

	for {
		var item entities.ItemPedido
		var produto *entities.Produto

		for {
			var nomeProduto string

			fmt.Printf("\nDigite o nome do produto: \n")
			fmt.Scanln(&nomeProduto)

			produto = BuscarProduto(nomeProduto, *produtosEmEstoque)
			if produto != nil {
				item.Nome = produto.Nome
				break
			}
			fmt.Println("\nProduto não encontrado. Tente novamente.")
		}

		for {
			var quantidade uint
			fmt.Print("\nDigite a quantidade: ")
			_, err := fmt.Scanln(&quantidade)

			if err == nil && quantidade > 0 {
				if quantidade <= produto.Quantidade {
					produto.Quantidade -= quantidade
					item.Quantidade = quantidade
					break
				} else {
					fmt.Printf("\nEstoque insuficiente.\n")
				}
			} else {
				fmt.Printf("\nApenas valores inteiros positivos serão aceitos.\n")
			}
		}

		item.Total = float64(item.Quantidade) * produto.Preco
		totalDoPedido += item.Total
		itensPedido = append(itensPedido, item)

		for {
			fmt.Println("\nCadastrar mais produtos?")
			fmt.Println("1: Sim")
			fmt.Println("2: Não")

			var opcao int
			_, err := fmt.Scanln(&opcao)

			if err == nil {
				if opcao == 2 {
					// Finaliza o pedido
					pedido := entities.Pedido{
						Produtos:      itensPedido,
						TotalDoPedido: totalDoPedido,
						Cliente:       clienteEmail,
					}
					*listaPedidos = append(*listaPedidos, pedido)
					entities.DadosDoPedido(pedido)
					return &pedido
				} else if opcao == 1 {
					break
				} else {
					fmt.Println("\nOpção inválida. Tente novamente.")
				}
			} else {
				fmt.Println("\nApenas valores inteiros serão aceitos.")
			}
		}
	}
}
