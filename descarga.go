package models

import "time"

type Descarga struct {
	ID        int       `json:"id"`
	UsuarioID int       `json:"usuario_id"`
	LibrosID  []int     `json:"libros_id"`
	Fecha     time.Time `json:"fecha"`
}
