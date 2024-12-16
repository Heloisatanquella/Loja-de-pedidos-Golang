package entities

import "fmt"

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade uint
}

func DadosProduto(p Produto) {
	fmt.Printf("\nNome do produto: %s; Dsecrição: %s; Preço unitário: %.2f; Quantidade disponível: %d;\n", p.Nome, p.Descricao, p.Preco, p.Quantidade)
}
