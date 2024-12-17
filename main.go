package main

import (
	"fmt"
	"loja-de-pedidos-golang/controllers"
	"loja-de-pedidos-golang/entities"
)

func main() {
	clientes := []entities.Cliente{
		{Email: "teste@exemplo.com"},
		{Email: "cliente@exemplo.com"},
	}

	fmt.Println(controllers.BuscarCliente("teste@exemplo.com",
		clientes))

	var clientesLista []entities.Cliente
	controllers.CadastrarCliente(&clientesLista)

}
