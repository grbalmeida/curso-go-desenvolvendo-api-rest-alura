package main

import (
	"github.com/grbalmeida/curso-go-desenvolvendo-api-rest-alura/database"
	"github.com/grbalmeida/curso-go-desenvolvendo-api-rest-alura/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequest()
}
