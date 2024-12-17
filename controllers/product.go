package controllers

import (
	"fmt"
	"loja-de-pedidos-golang/entities"
	"strings"
)

func BuscarProduto(nomeProduto string, produtosEmEstoque []entities.Produto) *entities.Produto {
	nomeProduto = strings.ToLower(strings.TrimSpace(nomeProduto))

	for _, produto := range produtosEmEstoque {
		if strings.ToLower(strings.TrimSpace(produto.Nome)) == nomeProduto {
			return &produto
		}
	}

	return nil
}

func Produto(produtosEmEstoque *[]entities.Produto) {
	var nomeProduto string
	var descricao string
	var preco float64
	var quantidade uint

	for {
		fmt.Printf("\nDigite o nome do produto: ")
		fmt.Scanln(&nomeProduto)

		if BuscarProduto(nomeProduto, *produtosEmEstoque) != nil {
			fmt.Println("\nProduto já existente. Tente novamente.")
		} else {
			break
		}
	}

	fmt.Printf("\nInsira uma descrição para o produto: ")
	fmt.Scanln(&descricao)

	for {
		fmt.Printf("\nInsira o preço do produto (00.00): R$ ")
		_, err := fmt.Scanln(&preco)

		if err != nil || preco <= 0 {
			fmt.Println("\nO valor inserido não é válido! Tente novamente.")
			var limpar string
			fmt.Scanln(&limpar)
			continue
		}
		break
	}

	for {
		fmt.Printf("\nInsira a quantidade em estoque: ")
		_, err := fmt.Scanln(&quantidade)

		if err != nil || quantidade <= 0 {
			fmt.Println("\nEntrada inválida. Por favor, insira um número inteiro positivo.")
			var limpar string
			fmt.Scanln(&limpar)
			continue
		}
		break
	}

	novoProduto := entities.Produto{
		Nome:       nomeProduto,
		Descricao:  descricao,
		Preco:      preco,
		Quantidade: quantidade,
	}

	*produtosEmEstoque = append(*produtosEmEstoque, novoProduto)

	fmt.Println("\nDados do produto cadastrado:")
	novoProduto.DadosProduto()
}
