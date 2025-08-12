package cli

import (
	"Cardapio-Digital/internal/models"
	"Cardapio-Digital/internal/repository"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MenuIngrediente(repo repository.IngredienteRepository) {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("INGREDIENTES")
		fmt.Println()
		fmt.Println("1 - LISTAR TODOS OS INGREDIENTES")
		fmt.Println("2 - INSERIR NOVO INGREDIENTE")
		fmt.Println("3 - ATUALIZAR INGREDIENTE")
		fmt.Println("4 - DELETAR INGREDIENTE")
		fmt.Println("0 - SAIR")
		fmt.Println()
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
			if len(nomeIngr) < 3 {
				fmt.Println("Nome muito curto, favor inserir nome com mais de três caracteres")
				continue
			}

			fmt.Println("Quantidade: ")
			quantStr, _ := reader.ReadString('\n')
			quantStr = strings.TrimSpace(quantStr)
			qntdIngr, err := strconv.Atoi(quantStr)
			q, err := strconv.Atoi(quantStr)
			if err != nil {
				fmt.Println("Quantidade inválida.")
				continue
			}
			if q < 0 {
				fmt.Println("Quantidade não pode ser menor que zero. Tente novamente.")
				continue
			}
			qntdIngr = q

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
			}
		case 3:
			var idIngrediente int
			fmt.Println("Digite o Id do Ingrediente: ")
			fmt.Scan(&idIngrediente)

			idIngredienteIdentificado, err := repo.IngredientePorId(idIngrediente)
			if err != nil {
				fmt.Println("Ingrediente nao encontrado", err)
			} else {
				jsonBytes, _ := json.MarshalIndent(idIngredienteIdentificado, "", "  ")
				fmt.Println(string(jsonBytes))
			}
			fmt.Println("Deseja atualizar o ingrediente ? (S / N)")
			var desejaAtualizar string
			fmt.Scan(&desejaAtualizar)
			if desejaAtualizar == "s" || desejaAtualizar == "S" {
				fmt.Println("Nome do Ingrediente: ")
				atualizaIngr, _ := reader.ReadString('\n')
				atualizaIngr = strings.TrimSpace(atualizaIngr)
				if len(atualizaIngr) < 3 {
					fmt.Println("Nome muito curto, favor inserir nome com mais de três caracteres")
					continue
				}

				novoIng := models.Ingrediente{
					Id:   idIngrediente,
					Nome: atualizaIngr,
				}
				if err := repo.AtualizarIngrediente(novoIng); err != nil {
					fmt.Println("Erro ao atualizar:", err)
				}
			} else {
				fmt.Println("Voltando para o menu inicial...")
				MenuIngrediente(repo)
			}
		case 4:
			var deletarIngr int
			fmt.Println("Digite o Id do Ingrediente: ")
			fmt.Scan(&deletarIngr)

			idIngredienteIdentificado, err := repo.IngredientePorId(deletarIngr)
			if err != nil {
				fmt.Println("Ingrediente nao encontrado", err)
			} else {
				jsonBytes, _ := json.MarshalIndent(idIngredienteIdentificado, "", "  ")
				fmt.Println(string(jsonBytes))
			}
			fmt.Println("Deseja deletar o Ingrediente desejado ? (S / N)")
			var opcaoDeletar string
			fmt.Scan(&opcaoDeletar)
			if opcaoDeletar == "s" || opcaoDeletar == "S" {
				delIng := models.Ingrediente{
					Id: deletarIngr,
				}
				if err := repo.DeletarIngrediente(delIng); err != nil {
					fmt.Println("Erro ao deletar:", err)
				}
			}
		case 0:
			fmt.Println("SAINDO...")
			os.Exit(0)
		}
	}
}
