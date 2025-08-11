package cli

import (
	"Cardapio-Digital/internal/database"
	"Cardapio-Digital/internal/models"
	"Cardapio-Digital/internal/repository"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Menu() {
	dsn := "root:senhasecreta@tcp(127.0.0.1:3306)/restaurante"
	db, err := database.NewMySQLDB(dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	repo := repository.IngredienteRepository{DB: db}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("INGREDIENTES")
		fmt.Println()
		fmt.Println("1 - LISTAR TODOS OS INGREDIENTES")
		fmt.Println("2 - INSERIR NOVO INGREDIENTE")
		fmt.Println("0 - SAIR")
		fmt.Print("Escolha uma opção: ")

		opcaoStr, _ := reader.ReadString('\n')
		opcaoStr = strings.TrimSpace(opcaoStr)
		opcao, err := strconv.Atoi(opcaoStr)
		if err != nil {
			fmt.Println("Opção inválida, tente novamente.")
			continue
		}

		switch opcao {
		case 1:
			allIngri, err := repo.TodosIngredientes()
			if err != nil {
				fmt.Println("Erro ao buscar ingredientes:", err)
				continue
			}
			jsonBytes, _ := json.MarshalIndent(allIngri, "", "  ")
			fmt.Println(string(jsonBytes))
		case 2:
			fmt.Println("Nome do Ingrediente: ")
			nomeIngr, _ := reader.ReadString('\n')
			nomeIngr = strings.TrimSpace(nomeIngr)

			fmt.Println("Quantidade: ")
			quantStr, _ := reader.ReadString('\n')
			quantStr = strings.TrimSpace(quantStr)
			qntdIngr, err := strconv.Atoi(quantStr)
			if err != nil {
				fmt.Println("Quantidade inválida.")
				continue
			}

			fmt.Println("Descrição: ")
			descrIngr, _ := reader.ReadString('\n')
			descrIngr = strings.TrimSpace(descrIngr)

			novoIng := models.Ingrediente{
				Nome:       nomeIngr,
				Quantidade: qntdIngr,
				Descricao:  descrIngr,
			}

			if err := repo.InserirIngrediente(novoIng); err != nil {
				fmt.Println("Erro ao inserir:", err)
			} else {
				fmt.Println("Ingrediente inserido com sucesso!")
			}
		case 0:
			fmt.Println("SAINDO...")
			os.Exit(0)
		}

	}
}
