package repository

import (
	"Cardapio-Digital/internal/models"
	"database/sql"
	"fmt"
)

type IngredienteRepository struct {
	DB *sql.DB
}

func (repo *IngredienteRepository) IngredientePorId(id int) (*models.Ingrediente, error) {
	var ingri models.Ingrediente
	err := repo.DB.QueryRow("SELECT * FROM INGREDIENTE WHERE ID = ?", id).
		Scan(&ingri.Id, &ingri.Nome, &ingri.Quantidade, &ingri.Descricao)
	if err != nil {
		return nil, err
	}
	return &ingri, nil
}

func (repo *IngredienteRepository) TodosIngredientes() ([]*models.Ingrediente, error) {
	res, err := repo.DB.Query("SELECT * FROM INGREDIENTE")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	ingredientes := []*models.Ingrediente{}
	for res.Next() {
		var ingri models.Ingrediente
		if err := res.Scan(&ingri.Id, &ingri.Nome, &ingri.Quantidade, &ingri.Descricao); err != nil {
			return nil, err
		}
		ingredientes = append(ingredientes, &ingri)
	}
	return ingredientes, nil
}

func (repo *IngredienteRepository) InserirIngrediente(ingri models.Ingrediente) error {
	_, err := repo.DB.Exec(
		"INSERT INTO INGREDIENTE (Nome, Quantidade, Descricao) VALUES (?, ?, ?)",
		ingri.Nome,
		ingri.Quantidade,
		ingri.Descricao,
	)
	if err != nil {
		return err
	}
	fmt.Println("------------------------------------------------------")
	fmt.Printf("Ingrediente %s inserido com sucesso\n", ingri.Nome)
	fmt.Println("------------------------------------------------------")
	return nil
}

func (repo *IngredienteRepository) AtualizarIngrediente(ingri models.Ingrediente) error {
	_, err := repo.DB.Exec(
		"UPDATE INGREDIENTE SET Nome = ? WHERE id = ?",
		ingri.Nome,
		ingri.Id,
	)
	if err != nil {
		return err
	}
	fmt.Println("------------------------------------------------------")
	fmt.Printf("Ingrediente %s foi alterado com sucesso!\n", ingri.Nome)
	fmt.Println("------------------------------------------------------")
	return nil
}

func (repo *IngredienteRepository) DeletarIngrediente(ingri models.Ingrediente) error {
	_, err := repo.DB.Exec(
		"DELETE FROM INGREDIENTE WHERE id = ?",
		ingri.Id,
	)
	if err != nil {
		return err
	}
	fmt.Println("------------------------------------------------------")
	fmt.Printf("Ingrediente %s foi excluido\n", ingri.Nome)
	fmt.Println("------------------------------------------------------")
	return nil
}
