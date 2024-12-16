package entities

import (
	"fmt"
)

type Client struct {
	Name      string `json:"name"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
	Telefone  string `json:"telefone"`
}

func DadosCompletos() {
	var client Client
	fmt.Printf("Nome completo: %s %s; E-mail: %s; Telefone: %s;\n", client.Name, client.Sobrenome, client.Email, client.Telefone)
}
