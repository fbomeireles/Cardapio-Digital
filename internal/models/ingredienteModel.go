package models

type Ingrediente struct {
	Id         int
	Nome       string
	Quantidade int
	Descricao  string
}
type Prato struct {
	Id        int
	Nome      string
	Descricao string
	URL_Foto  string
}
type Prato_Ingrediente struct {
	Id_Prato       int
	Id_Ingrediente int
	Quantidade     int
}
