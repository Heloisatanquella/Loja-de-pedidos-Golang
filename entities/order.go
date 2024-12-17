package entities

import "fmt"

type ItemPedido struct {
	Nome       string
	Quantidade uint
	Total      float64
}

type Pedido struct {
	Produtos      []ItemPedido
	TotalDoPedido float64
	Cliente       string
}

func (i ItemPedido) DadosItem() {
	fmt.Printf("Item: %s; Quantidade: %d; Total: R$%.2f;\n", i.Nome, i.Quantidade, i.Total)
}

func DadosDoPedido(p Pedido) {
	fmt.Println("\nProdutos adicionados ao pedido:")
	for _, item := range p.Produtos {
		item.DadosItem()
	}
	fmt.Printf("\nValor total do pedido: R$%.2f; Comprador: %s;\n", p.TotalDoPedido, p.Cliente)
}
