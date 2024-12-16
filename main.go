package main

import (
	"loja-de-pedidos-golang/controllers"
	"loja-de-pedidos-golang/entities"
)

func main() {
	cliente1 := entities.Cliente{Nome: "João Silva", Email: "joao@exemplo.com", Telefone: "47 99266-7703"}
	cliente2 := entities.Cliente{Nome: "Maria Helena", Email: "maria@exemplo.com", Telefone: "48 99156-1809"}

	clientes := []entities.Cliente{cliente1, cliente2}

	controllers.BuscarCliente(clientes)
}
