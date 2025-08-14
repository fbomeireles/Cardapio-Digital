package handlers

import (
	"Cardapio-Digital/internal/models"
	"Cardapio-Digital/internal/repository"
	"encoding/json"
	"net/http"
)

func CriarPrato(repo *repository.PratoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var prato models.Prato
		if err := json.NewDecoder(r.Body).Decode(&prato); err != nil {
			http.Error(w, "JSON Inválido", http.StatusBadRequest)
			return
		}
		if prato.Nome == "" {
			http.Error(w, "Campo nome é obrigatório", http.StatusBadRequest)
			return
		}
		ok, _ := repo.PratoPorNome(prato.Nome)
		if ok != nil {
			http.Error(w, "Prato já existe", http.StatusConflict)
			return
		}
		if err := repo.InserirPrato(prato); err != nil {
			http.Error(w, "Erro ao cadastrar", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Ingrediente cadastrado!"))
	}
}
