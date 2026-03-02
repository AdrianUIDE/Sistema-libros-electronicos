package models

type Libro struct {
	ID         int    `json:"id"`
	Titulo     string `json:"titulo"`
	Autor      string `json:"autor"`
	ISBN       string `json:"isbn"`
	Descargado bool   `json:"descargado"`
}
