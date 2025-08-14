package models

type Prato_Ingrediente struct {
	Id_Prato       int `json:"Id_Prato"`
	Id_Ingrediente int `json:"Id_Ingrediente"`
	Quantidade     int `json:"Quantidade"`
}
