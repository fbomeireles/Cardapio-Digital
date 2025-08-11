package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Ingrediente struct {
	Id         int
	Nome       string
	Quantidade int
	Descricao  string
}
type Prato struct {
	Id        int
	Nome      string
	Descricao string
	URL_Foto  string
}
type Prato_Ingrediente struct {
	Id_Prato       int
	Id_Ingrediente int
	Quantidade     int
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:senhasecreta@tcp(127.0.0.1:3306)/restaurante")
	if err != nil {
		fmt.Println("Não foi possível estabelecer conexão com o banco de dados. Por favor, verifique as configurações e tente novamente.")
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Não foi possível conectar ao banco de dados (ping).", err)
		os.Exit(1)
	}
	fmt.Println("======================================================")
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")
	fmt.Println("======================================================")
	fmt.Println("")

	// fmt.Println("INSERIR INSERIR INSERIR INSERIR INSERIR")
	// ingri := Ingrediente{
	// 	Nome:       "Alhos",
	// 	Quantidade: 99,
	// 	Descricao:  "",
	// }
	// if insertError := InserirIngrediente(ingri); insertError != nil {
	// 	fmt.Println("Erro ao inserir o ingrediente:", insertError)
	// 	os.Exit(1)
	// }
	// fmt.Println("")
	fmt.Println("ATUALIZAR ATUALIZAR ATUALIZAR ATUALIZAR ATUALIZAR")
	atualizaIngr := Ingrediente{
		Id:   4,
		Nome: "Alho",
	}
	if atualizaErr := AtualizarIngrediente(atualizaIngr); atualizaErr != nil {
		fmt.Println("Erro ao atualizar o ingrediente: ", atualizaErr)
	}
	fmt.Println("")
	// fmt.Println("DELETAR DELETAR DELETAR DELETAR DELETAR")
	// delIngrediente := Ingrediente{
	// 	Id: 3,
	// }
	// errDelIngr := DeletarIngrediente(delIngrediente)
	// if errDelIngr != nil {
	// 	fmt.Println("Erro ao deletar ingrediente", errDelIngr)
	// 	os.Exit(1)
	// }
	// fmt.Println("")
	fmt.Println("BUSCAR BUSCAR BUSCAR BUSCAR BUSCAR")
	allIngri, errAll := TodosIngredientes()
	if errAll != nil {
		fmt.Println("Erro ao buscar ingredientes:", errAll)
		os.Exit(1)
	}
	jsonBytes, _ := json.MarshalIndent(allIngri, "", "  ")
	fmt.Println(string(jsonBytes))
	fmt.Println(allIngri)
	fmt.Println("")
}
func TodosIngredientes() ([]*Ingrediente, error) {
	res, err := db.Query("SELECT * FROM INGREDIENTE")

	if err != nil {
		return nil, err
	}

	ingredientes := []*Ingrediente{}

	for res.Next() {
		var ingri Ingrediente

		if err := res.Scan(&ingri.Id, &ingri.Nome, &ingri.Quantidade, &ingri.Descricao); err != nil {
			return nil, err
		}
		ingredientes = append(ingredientes, &ingri)
	}
	return ingredientes, nil
}

func InserirIngrediente(ingredientes Ingrediente) error {
	_, err := db.Exec(
		"INSERT INTO INGREDIENTE (Nome, Quantidade, Descricao) VALUES (?, ?, ?)",
		ingredientes.Nome,
		ingredientes.Quantidade,
		ingredientes.Descricao,
	)
	if err != nil {
		return err
	}
	fmt.Println("")
	fmt.Println("------------------------------------------------------")
	fmt.Println("Ingrediente", ingredientes.Nome, "inserido com sucesso")
	fmt.Println("------------------------------------------------------")
	fmt.Println("")
	return nil
}
func AtualizarIngrediente(ingredientes Ingrediente) error {
	_, err := db.Exec(
		"UPDATE INGREDIENTE SET Nome = ? WHERE id = ?",
		ingredientes.Nome,
		ingredientes.Id,
	)
	if err != nil {
		return err
	}
	fmt.Println("")
	fmt.Println("------------------------------------------------------")
	fmt.Println("Ingrediente", ingredientes.Nome, "foi alterado com sucesso!")
	fmt.Println("------------------------------------------------------")
	fmt.Println("")
	return nil
}
func DeletarIngrediente(ingredientes Ingrediente) error {
	_, err := db.Exec(
		"DELETE FROM INGREDIENTE WHERE id = ?",
		ingredientes.Id,
	)
	if err != nil {
		return err
	}
	fmt.Println("")
	fmt.Println("------------------------------------------------------")
	fmt.Println("Ingrediente", ingredientes.Nome, "foi excluido")
	fmt.Println("------------------------------------------------------")
	fmt.Println("")
	return nil
}
