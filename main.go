package main

import (
	"fmt"

	"github.com/betocalestini/api_rest_go/database"
	"github.com/betocalestini/api_rest_go/models"
	"github.com/betocalestini/api_rest_go/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nome 1", Historia: "historia 1"},
		{Id: 2, Nome: "nome 2", Historia: "historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Servidor go rodando")
	routes.HandleRequest()
}
