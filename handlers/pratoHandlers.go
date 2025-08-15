package handlers

import (
	"Cardapio-Digital/internal/models"
	"Cardapio-Digital/internal/repository"
	"encoding/json"
	"net/http"
)

type IngredienteSolicit struct {
	IdIngrediente int `json:"id_ingrediente"`
	Quantidade    int `json:"quantidade"`
}
type PratoSolicit struct {
	Nome         string               `json:"nome"`
	Descricao    *string              `json:"descricao,omitempty"`
	UrlFoto      *string              `json:"url_foto,omitempty"`
	Ingredientes []IngredienteSolicit `json:"ingredientes"`
}

func CriarPrato(repo *repository.PratoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req PratoSolicit
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		if len(req.Nome) < 3 {
			http.Error(w, "Nome obrigatório (mínimo 3 letras)", http.StatusBadRequest)
			return
		}

		if len(req.Ingredientes) == 0 {
			http.Error(w, "A lista de ingredientes está vazia", http.StatusBadRequest)
			return
		}

		for _, ing := range req.Ingredientes {
			if ing.Quantidade == 0 {
				http.Error(w, "Ingrediente não pode ter quantidade zero", http.StatusBadRequest)
				return
			}
		}

		foto := "https://metanikk.com.br"
		if req.UrlFoto != nil && *req.UrlFoto != "" {
			foto = *req.UrlFoto
		}

		existe, _ := repo.PratoPorNome(req.Nome)
		if existe != nil {
			http.Error(w, "Prato já existe", http.StatusConflict)
			return
		}

		prato := models.Prato{
			Nome:      req.Nome,
			Descricao: "",
			URL_Foto:  foto,
		}
		if req.Descricao != nil {
			prato.Descricao = *req.Descricao
		}

		// 8. Montar a lista de ingredientes para salvar
		ingredientes := make([]models.Prato_Ingrediente, len(req.Ingredientes))
		for i, ing := range req.Ingredientes {
			ingredientes[i] = models.Prato_Ingrediente{
				Id_Ingrediente: ing.IdIngrediente,
				Quantidade:     ing.Quantidade,
			}
		}

		err := repo.InserirPrato(prato)
		if err != nil {
			http.Error(w, "Erro ao cadastrar prato", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Prato cadastrado com sucesso!"))
	}
}
