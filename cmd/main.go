package main

import (
	"Cardapio-Digital/internal/cli"
	"Cardapio-Digital/internal/database"
	"Cardapio-Digital/internal/repository"
)

func main() {
	dsn := "root:senhasecreta@tcp(127.0.0.1:3306)/restaurante"

	db, err := database.NewMySQLDB(dsn)
	if err != nil {
		panic(err)
	}

	ingredienteRepo := repository.IngredienteRepository{DB: db}
	pratoRepo := repository.PratoRepository{DB: db}

	cli.MenuIngrediente(ingredienteRepo)
	cli.MenuPrato(pratoRepo)

}
