package main

import (
	"Cardapio-Digital/handlers"
	"Cardapio-Digital/internal/database"
	"Cardapio-Digital/internal/repository"
	"fmt"
	"net/http"
	"os"
)

func main() {

	dsn := "root:senhasecreta@tcp(127.0.0.1:3306)/restaurante"
	db, err := database.NewMySQLDB(dsn)
	if err != nil {
		fmt.Println("Não foi possível conectar ao banco")
		os.Exit(1)
	}
	repoIngr := &repository.IngredienteRepository{DB: db}

	http.HandleFunc("/ingredientes/criar", handlers.CriarIngrediente(repoIngr))
	http.HandleFunc("/ingredientes/editar", handlers.EditarIngrediente(repoIngr))
	http.HandleFunc("/ingredientes/inativar/", handlers.InativarIngrediente(repoIngr))
	http.HandleFunc("/ingredientes/todos", handlers.BuscarTodosOsIngredientes(repoIngr))
	http.HandleFunc("/ingredientes/", handlers.BuscarIngredientePorId(repoIngr))

	fmt.Println("Servidor rodando em http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

	// cli.Menu()

}
