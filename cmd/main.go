package main

import (
	"Cardapio-Digital/internal/cli"
)

func main() {
	// dsn := "root:senhasecreta@tcp(127.0.0.1:3306)/restaurante"
	// db, err := database.NewMySQLDB(dsn)
	// if err != nil {
	// 	fmt.Println(err)

	// 	os.Exit(1)
	// }
	// repo := repository.IngredienteRepository{DB: db}
	cli.Menu()

	// Deletar exemplo
	// delIng := models.Ingrediente{
	// 	Id: 3,
	// }
	// if err := repo.DeletarIngrediente(delIng); err != nil {
	// 	fmt.Println("Erro ao deletar:", err)
	// }

	// Buscar todos

}
