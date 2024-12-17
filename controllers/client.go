package controllers

import (
	"bufio"
	"fmt"
	"loja-de-pedidos-golang/entities"
	"loja-de-pedidos-golang/utils"
	"os"
	"strings"
)

func BuscarCliente(email string, clientesLista []entities.Cliente) bool {
	email = strings.ToLower(strings.TrimSpace(email))

	return utils.Buscar(clientesLista, func(cliente entities.Cliente) bool {
		return strings.ToLower(strings.TrimSpace(cliente.Email)) == email
	})
}

func CadastrarCliente(clientesLista *[]entities.Cliente) entities.Cliente {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nDigite o primeiro nome do cliente: ")
	nome, _ := reader.ReadString('\n')
	nome = strings.TrimSpace(nome)

	fmt.Print("\nDigite o sobrenome do cliente: ")
	sobrenome, _ := reader.ReadString('\n')
	sobrenome = strings.TrimSpace(sobrenome)

	var email string
	for {
		fmt.Print("\nInsira o e-mail do cliente: ")
		email, _ = reader.ReadString('\n')
		email = strings.TrimSpace(email)

		if BuscarCliente(email, *clientesLista) {
			fmt.Printf("\nE-mail já cadastrado. Tente novamente.\n")
		} else {
			break
		}
	}

	fmt.Print("\nInsira o número para contato do cliente (DDD 9 9999.9999): ")
	telefone, _ := reader.ReadString('\n')
	telefone = strings.TrimSpace(telefone)

	novoCliente := entities.Cliente{
		Nome:      nome,
		Sobrenome: sobrenome,
		Email:     email,
		Telefone:  telefone,
	}

	*clientesLista = append(*clientesLista, novoCliente)

	fmt.Println("\nDados do cliente cadastrado:")
	novoCliente.DadosCompletos()

	return novoCliente
}