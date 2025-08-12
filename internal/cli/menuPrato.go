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

func MenuPrato(repo repository.PratoRepository) {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("PRATOS")
		fmt.Println()
		fmt.Println("1 - LISTAR TODOS OS PRATOS")
		fmt.Println("2 - INSERIR NOVO PRATOS")
		fmt.Println("3 - ATUALIZAR PRATOS")
		fmt.Println("4 - DELETAR PRATOS")
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
			allIngri, err := repo.TodosOsPratos()
			if err != nil {
				fmt.Println("Erro ao buscar os pratos:", err)
				continue
			}
			jsonBytes, _ := json.MarshalIndent(allIngri, "", "  ")
			fmt.Println(string(jsonBytes))
		case 2:
			fmt.Println("Nome do Prato: ")
			nomePrato, _ := reader.ReadString('\n')
			nomePrato = strings.TrimSpace(nomePrato)
			if len(nomePrato) < 3 {
				fmt.Println("Nome muito curto, favor inserir nome com mais de três caracteres")
				continue
			}
			fmt.Println("Descrição: ")
			descrPrato, _ := reader.ReadString('\n')
			descrPrato = strings.TrimSpace(descrPrato)

			fmt.Println("URL Foto: ")
			urlPrato, _ := reader.ReadString('\n')
			urlPrato = strings.TrimSpace(urlPrato)

			novoPrato := models.Prato{
				Nome:      nomePrato,
				Descricao: descrPrato,
				URL_Foto:  urlPrato,
			}

			if err := repo.InserirPrato(novoPrato); err != nil {
				fmt.Println("Erro ao inserir:", err)
			}
		case 3:
			var idPrato int
			fmt.Println("Digite o Id do Prato: ")
			fmt.Scan(&idPrato)

			idIPratoIdentificado, err := repo.PratoPorId(idPrato)
			if err != nil {
				fmt.Println("Prato nao encontrado", err)
			} else {
				jsonBytes, _ := json.MarshalIndent(idIPratoIdentificado, "", "  ")
				fmt.Println(string(jsonBytes))
			}
			fmt.Println("Deseja atualizar o Prato ? (S / N)")
			var desejaAtualizar string
			fmt.Scan(&desejaAtualizar)
			if desejaAtualizar == "s" || desejaAtualizar == "S" {
				fmt.Println("Nome do Prato: ")
				atualizaPrato, _ := reader.ReadString('\n')
				atualizaPrato = strings.TrimSpace(atualizaPrato)
				if len(atualizaPrato) < 3 {
					fmt.Println("Nome muito curto, favor inserir nome com mais de três caracteres")
					continue
				}

				novoPrato := models.Prato{
					Id:   idPrato,
					Nome: atualizaPrato,
				}
				if err := repo.AtualizarPrato(novoPrato); err != nil {
					fmt.Println("Erro ao atualizar:", err)
				}
			} else {
				fmt.Println("Voltando para o menu inicial...")
				MenuPrato(repo)
			}
		case 4:
			var deletarPrato int
			fmt.Println("Digite o Id do Ingrediente: ")
			fmt.Scan(&deletarPrato)

			idPratoIdentificado, err := repo.PratoPorId(deletarPrato)
			if err != nil {
				fmt.Println("Ingrediente nao encontrado", err)
			} else {
				jsonBytes, _ := json.MarshalIndent(idPratoIdentificado, "", "  ")
				fmt.Println(string(jsonBytes))
			}
			fmt.Println("Deseja deletar o Ingrediente desejado ? (S / N)")
			var opcaoDeletar string
			fmt.Scan(&opcaoDeletar)
			if opcaoDeletar == "s" || opcaoDeletar == "S" {
				delPrato := models.Prato{
					Id: deletarPrato,
				}
				if err := repo.DeletarPrato(delPrato); err != nil {
					fmt.Println("Erro ao deletar:", err)
				}
			}
		case 0:
			fmt.Println("SAINDO...")
			os.Exit(0)
		}
	}
}
