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
	db, err = sql.Open("mysql", "root:senhasecreta@tcp(127.0.0.1.3306)/restaurante")

	if err != nil {
		fmt.Println("Não foi possível estabelecer conexão com o banco de dados. Por favor, verifique as configurações e tente novamente.")
		os.Exit(1)
	}
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")
}
