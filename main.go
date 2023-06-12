package main

import (
	"fmt"

	"github.com/thislucio/go-alura-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
