package entities

import "fmt"

type ItemPedido struct {
	Nome       string  `json:"nome"`
	Quantidade uint    `json:"quantidade"`
	Total      float64 `json:"total"`
}

type Pedido struct {
	Produtos      []interface{} `json:"produtos"`
	TotalDoPedido float64       `json:"totaldopedido"`
	Cliente       string        `json:"cliente"`
}

func (i ItemPedido) DadosItem() {
	fmt.Printf("Item: %s; Quantidade: %d; Total: R$%.2f;\n", i.Nome, i.Quantidade, i.Total)
}

func DadosDoPedido(p Pedido) {
	fmt.Printf("Produtos adicionados ao pedido: ")
	for _, produto := range p.Produtos {
		if item, ok := produto.(ItemPedido); ok {
			item.DadosItem()
		}
	}
	fmt.Printf("\nValor total do pedido: R$%.2f; Comprador: %s\n", p.TotalDoPedido, p.Cliente)
}
