package main

import (
	"fmt"
	"net/http"

	"evaluacionfinal/config"
	"evaluacionfinal/handlers"
)

func main() {

	// CONECTAR A BASE DE DATOS
	config.Connect()

	http.HandleFunc("/usuarios", handlers.CrearUsuario)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
