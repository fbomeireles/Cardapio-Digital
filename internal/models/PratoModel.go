package models

type Prato struct {
	Id        int    `json:"Id"`
	Nome      string `json:"Nome"`
	Descricao string `json:"Descricao"`
	URL_Foto  string `json:"URL_Foto"`
}
