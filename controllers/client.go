package controllers

import (
	"fmt"
	"loja-de-pedidos-golang/entities"
	"strings"
)

func BuscarCliente(clientes []entities.Cliente) *entities.Cliente {
	var email string
	for {
		fmt.Print("\nDigite o e-mail do cliente: ")
		fmt.Scanln(&email)

		for i := range clientes {
			if strings.ToLower(clientes[i].Email) == strings.ToLower(email) {
				clientes[i].DadosCompletos()
				return &clientes[i]
			}
		}

		fmt.Println("\nCliente não encontrado!")
	}
}
