package main

import (
	"Cardapio-Digital/internal/database"
	"Cardapio-Digital/internal/models"
	"Cardapio-Digital/internal/repository"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	dsn := "root:senhasecreta@tcp(127.0.0.1:3306)/restaurante"
	db, err := database.NewMySQLDB(dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	repo := repository.IngredienteRepository{DB: db}

	// Inserir exemplo
	// novoIng := models.Ingrediente{
	// 	Nome:       "Alho",
	// 	Quantidade: 99,
	// 	Descricao:  "Ingrediente usado em muitos pratos",
	// }
	// if err := repo.InserirIngrediente(novoIng); err != nil {
	// 	fmt.Println("Erro ao inserir:", err)
	// }

	// Atualizar exemplo
	fmt.Println("ATUALIZAR")
	atualizaIng := models.Ingrediente{
		Id:   4,
		Nome: "Alho Atualizado",
	}
	if err := repo.AtualizarIngrediente(atualizaIng); err != nil {
		fmt.Println("Erro ao atualizar:", err)
	}

	// Deletar exemplo
	// delIng := models.Ingrediente{
	// 	Id: 3,
	// }
	// if err := repo.DeletarIngrediente(delIng); err != nil {
	// 	fmt.Println("Erro ao deletar:", err)
	// }

	// Buscar todos
	fmt.Println("BUSCAR")
	allIngri, err := repo.TodosIngredientes()
	if err != nil {
		fmt.Println("Erro ao buscar ingredientes:", err)
		os.Exit(1)
	}
	jsonBytes, _ := json.MarshalIndent(allIngri, "", "  ")
	fmt.Println(string(jsonBytes))
}
