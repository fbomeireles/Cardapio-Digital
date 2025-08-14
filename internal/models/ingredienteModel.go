package models

type Ingrediente struct {
	Id         int    `json:"Id"`
	Nome       string `json:"Nome"`
	Quantidade int    `json:"Quantidade"`
	Descricao  string `json:"Descricao"`
}
