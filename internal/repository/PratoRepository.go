package repository

import (
	"Cardapio-Digital/internal/models"
	"database/sql"
	"fmt"
)

type PratoRepository struct {
	DB *sql.DB
}

func (repo *PratoRepository) PratoPorId(id int) (*models.Prato, error) {
	var prato models.Prato
	err := repo.DB.QueryRow("SELECT * FROM PRATO WHERE ID = ?", id).
		Scan(&prato.Id, &prato.Nome, &prato.Descricao, &prato.URL_Foto)
	if err != nil {
		return nil, err
	}
	return &prato, nil
}
func (repo *PratoRepository) TodosOsPratos() ([]*models.Prato, error) {
	res, err := repo.DB.Query("SELECT * FROM PRATO")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	pratos := []*models.Prato{}
	for res.Next() {
		var prato models.Prato
		if err := res.Scan(&prato.Id, &prato.Nome, &prato.Descricao, &prato.URL_Foto); err != nil {
			return nil, err
		}
		pratos = append(pratos, &prato)
	}
	return pratos, nil
}
func (repo *PratoRepository) InserirPrato(prato models.Prato) error {
	_, err := repo.DB.Exec("INSERT INTO PRATO (Nome, Descricao, URL_Foto) VALUES ? ? ?",
		prato.Nome,
		prato.Descricao,
		prato.URL_Foto,
	)
	if err != nil {
		return err
	}
	fmt.Println("------------------------------------------------------")
	fmt.Printf("Prato %s inserido com sucesso\n", prato.Nome)
	fmt.Println("------------------------------------------------------")
	return nil
}
func (repo *PratoRepository) AtualizarPrato(prato models.Prato) error {
	_, err := repo.DB.Exec("UPDATE PRATO SET Nome = ? WHERE Id = ?",
		prato.Id,
		prato.Nome,
	)
	if err != nil {
		return err
	}
	fmt.Println("------------------------------------------------------")
	fmt.Printf("Prato %s foi alterado com sucesso!\n", prato.Nome)
	fmt.Println("------------------------------------------------------")
	return nil
}
func (repo *PratoRepository) DeletarPrato(prato models.Prato) error {
	_, err := repo.DB.Exec(
		"DELETE FROM PRATO WHERE id = ?",
		prato.Id,
	)
	if err != nil {
		return err
	}
	fmt.Println("------------------------------------------------------")
	fmt.Printf("Prato %s foi excluido\n", prato.Nome)
	fmt.Println("------------------------------------------------------")
	return nil
}
