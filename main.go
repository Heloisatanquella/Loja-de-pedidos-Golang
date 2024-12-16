package main

import (
	"fmt"
	"loja-de-pedidos-golang/entities"
)

func main() {
	fmt.Println("Loja de pedidos ok!")
	entities.DadosCompletos(entities.Cliente{})
	entities.DadosDoPedido(entities.Pedido{})
	entities.DadosProduto(entities.Produto{})
}
