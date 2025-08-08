package main

import (
	"database/sql"
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
	fmt.Println("======================================================")
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")
	fmt.Println("======================================================")

	ingri := Ingrediente{
		Nome:       "Cebola",
		Quantidade: 99,
		Descricao:  "",
	}
	if insertError := InserirIngrediente(ingri); insertError != nil {
		panic(insertError)
	}
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
