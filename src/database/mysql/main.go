package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_, err := sql.Open("mysql", "root:senhasecreta@tcp(127.0.0.1.3306)/fbomrl")

	if err != nil {
		fmt.Println("Não foi possível estabelecer conexão com o banco de dados. Por favor, verifique as configurações e tente novamente.")
		os.Exit(1)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")
}
