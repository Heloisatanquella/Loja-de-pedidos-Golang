package entities

import (
	"fmt"
)

type Cliente struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
	Telefone  string `json:"telefone"`
}

func DadosCompletos(c Cliente) {
	fmt.Printf("Nome completo: %s %s; E-mail: %s; Telefone: %s;\n", c.Nome, c.Sobrenome, c.Email, c.Telefone)
}
