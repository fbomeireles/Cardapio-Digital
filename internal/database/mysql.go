package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Erro ao abrir conexão: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Erro ao conectar (ping): %w", err)
	}
	fmt.Println("======================================================")
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")
	fmt.Println("======================================================")
	fmt.Println("")
	return db, nil
}
