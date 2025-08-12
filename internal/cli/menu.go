package cli

import (
	"Cardapio-Digital/internal/database"
	"Cardapio-Digital/internal/repository"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Menu() {
	dsn := "root:senhasecreta@tcp(127.0.0.1:3306)/restaurante"

	db, err := database.NewMySQLDB(dsn)
	if err != nil {
		panic(err)
	}

	ingredienteRepo := repository.IngredienteRepository{DB: db}
	pratoRepo := repository.PratoRepository{DB: db}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("SELECIONE UMA OPCAO: ")
		fmt.Println("")
		fmt.Println("1 - INGREDIENTES")
		fmt.Println("2 - PRATOS")
		fmt.Println("")
		opcaoStr, _ := reader.ReadString('\n')
		opcaoStr = strings.TrimSpace(opcaoStr)
		opcao, err := strconv.Atoi(opcaoStr)
		if err != nil {
			fmt.Println("Opção inválida, tente novamente.")
			continue
		}

		switch opcao {
		case 1:
			MenuIngrediente(ingredienteRepo)
		case 2:
			MenuPrato(pratoRepo)
		case 0:
			fmt.Println("SAINDO...")
			os.Exit(0)

		}
	}
}
