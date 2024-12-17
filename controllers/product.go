package controllers

import (
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
