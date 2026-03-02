package handlers

import (
	"encoding/json"
	"net/http"

	"evaluacionfinal/config"
	"evaluacionfinal/models"
)

func CrearUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	if usuario.Nombres == "" || usuario.Apellidos == "" || usuario.Email == "" {
		http.Error(w, "Campos obligatorios", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO usuarios (nombres, apellidos, email) VALUES (@p1, @p2, @p3)"
	_, err := config.DB.Exec(query, usuario.Nombres, usuario.Apellidos, usuario.Email)

	if err != nil {
		http.Error(w, "Error al insertar", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}
