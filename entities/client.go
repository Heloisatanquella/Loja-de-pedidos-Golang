package entities

import (
	"fmt"
)

type Cliente struct {
	Nome      string
	Sobrenome string
	Email     string
	Telefone  string
}

func (c *Cliente) DadosCompletos() {
	fmt.Printf("\nNome completo: %s %s; E-mail: %s; Telefone: %s;\n\n", c.Nome, c.Sobrenome, c.Email, c.Telefone)
}
