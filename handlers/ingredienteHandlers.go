package handlers

import (
	"Cardapio-Digital/internal/models"
	"Cardapio-Digital/internal/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func CriarIngrediente(repo *repository.IngredienteRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ingr models.Ingrediente
		if err := json.NewDecoder(r.Body).Decode(&ingr); err != nil {
			http.Error(w, "JSON Inválido", http.StatusBadRequest)
			return
		}
		if ingr.Nome == "" || len(ingr.Nome) < 3 || ingr.Quantidade < 0 {
			http.Error(w, "Campo obrigatório inválido", http.StatusBadRequest)
			return
		}
		ok, _ := repo.IngredientePorNome(ingr.Nome)
		if ok != nil {
			http.Error(w, "Ingrediente já existe!", http.StatusConflict)
			return
		}
		if err := repo.InserirIngrediente(ingr); err != nil {
			http.Error(w, "Erro ao cadastrar", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Ingrediente cadastrado!"))
	}
}
func EditarIngrediente(repo *repository.IngredienteRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ingr models.Ingrediente
		if err := json.NewDecoder(r.Body).Decode(&ingr); err != nil {
			http.Error(w, "JSON Inválido", http.StatusBadRequest)
			return
		}
		if ingr.Nome == "" || len(ingr.Nome) < 3 || ingr.Quantidade < 0 {
			http.Error(w, "Campo obrigatório inválido", http.StatusBadRequest)
			return
		}
		_, err := repo.IngredientePorId(ingr.Id)
		if err != nil {
			fmt.Printf("Falha na busca do id %d: %v\n", ingr.Id, err)
			http.Error(w, "Ingrediente não existe!", http.StatusNotFound)
			return
		}
		ingrExistente, err := repo.IngredientePorNome(ingr.Nome)
		if err == nil && ingrExistente.Id != ingr.Id {
			http.Error(w, "Ingrediente já existe.", http.StatusConflict)
			return
		}
		if err := repo.AtualizarIngrediente(ingr); err != nil {
			http.Error(w, "Erro ao atualizar", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Ingrediente alterado!"))
	}
}

// Verificar como Inativar depois...//
func InativarIngrediente(repo *repository.IngredienteRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ingr models.Ingrediente
		if err := json.NewDecoder(r.Body).Decode(&ingr); err != nil {
			http.Error(w, "JSON Inválido", http.StatusBadRequest)
			return
		}
		_, err := repo.IngredientePorId(ingr.Id)
		if err != nil {
			http.Error(w, "Ingrediente não existe na base de dados", http.StatusNotFound)
		}
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Ingrediente deletado!"))
	}
}

func BuscarIngredientePorId(repo *repository.IngredienteRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		url := r.URL.Path
		idString := url[len("/ingredientes/"):]

		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Id inválido", http.StatusBadRequest)
		}

		ingr, err := repo.IngredientePorId(id)
		if err != nil {
			http.Error(w, "Ingrediente não existe na base de dados", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ingr)
	}
}
func BuscarTodosOsIngredientes(repo *repository.IngredienteRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ingredientes, err := repo.TodosIngredientes()
		if err != nil {
			http.Error(w, "Erro ao buscar ingredientes", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ingredientes)
	}
}
