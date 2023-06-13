package main

import (
	"fmt"

	"github.com/thislucio/go-alura-rest-api/database"
	"github.com/thislucio/go-alura-rest-api/models"
	"github.com/thislucio/go-alura-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "NOme 2", Historia: "Historia 2"},
	}
	database.ConectaComDb()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
