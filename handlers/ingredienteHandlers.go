package handlers

import (
	"Cardapio-Digital/internal/models"
	"Cardapio-Digital/internal/repository"
	"encoding/json"
	"net/http"
)

func CriarIngredienteHandler(repo *repository.IngredienteRepository) http.HandlerFunc {
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
