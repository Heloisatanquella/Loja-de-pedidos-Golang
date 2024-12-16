package entities

import "fmt"

type Produto struct {
	Nome       string  `json:"nome"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco"`
	Quantidade uint    `json:"quantidade"`
}

func DadosProduto(p Produto) {
	fmt.Printf("\nNome do produto: %s; Dsecrição: %s; Preço unitário: %.2f; Quantidade disponível: %d;\n", p.Nome, p.Descricao, p.Preco, p.Quantidade)
}
