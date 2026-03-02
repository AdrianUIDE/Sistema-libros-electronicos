/*
@autor: Adrian Herrera
@fecha: 01/03/2026
@version: 1.0
@descripcion: Proyecto Final
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

// ======================
// ESTRUCTURAS
// ======================

type Usuario struct {
	ID        int    `json:"id"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Email     string `json:"email"`
}

type Libro struct {
	ID         int    `json:"id"`
	Titulo     string `json:"titulo"`
	Autor      string `json:"autor"`
	ISBN       string `json:"isbn"`
	Descargado bool   `json:"descargado"`
}

type Descarga struct {
	ID        int       `json:"id"`
	UsuarioID int       `json:"usuario_id"`
	LibrosID  []int     `json:"libros_id"`
	Fecha     time.Time `json:"fecha"`
}

// ======================
// VARIABLES GLOBALES
// ======================

var usuarios []Usuario
var libros []Libro
var descargas []Descarga
var ultimoID int
var mutex = &sync.Mutex{}

// ======================
// HANDLERS USUARIOS
// ======================

func crearUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	// VALIDACIONES
	if usuario.Nombres == "" || usuario.Apellidos == "" || usuario.Email == "" {
		http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
		return
	}

	if !strings.Contains(usuario.Email, "@") {
		http.Error(w, "Email inválido", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	ultimoID++
	usuario.ID = ultimoID
	usuarios = append(usuarios, usuario)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}

func listarUsuarios(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(usuarios)
}

// ======================
// HANDLERS LIBROS
// ======================

func crearLibro(w http.ResponseWriter, r *http.Request) {
	var libro Libro
	json.NewDecoder(r.Body).Decode(&libro)

	// VALIDACIONES
	if libro.Titulo == "" || libro.Autor == "" || libro.ISBN == "" {
		http.Error(w, "Todos los campos del libro son obligatorios", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	ultimoID++
	libro.ID = ultimoID
	libro.Descargado = false
	libros = append(libros, libro)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(libro)
}

func listarLibros(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(libros)
}

// ======================
// HANDLERS DESCARGAS
// ======================

func crearDescarga(w http.ResponseWriter, r *http.Request) {
	var descarga Descarga
	json.NewDecoder(r.Body).Decode(&descarga)

	// VALIDAR USUARIO
	usuarioExiste := false
	for _, u := range usuarios {
		if u.ID == descarga.UsuarioID {
			usuarioExiste = true
			break
		}
	}

	if !usuarioExiste {
		http.Error(w, "El usuario no existe", http.StatusBadRequest)
		return
	}

	// VALIDAR LIBROS
	for _, libroID := range descarga.LibrosID {
		libroExiste := false

		for i := range libros {
			if libros[i].ID == libroID {
				if libros[i].Descargado {
					http.Error(w, "El libro ya fue descargado", http.StatusBadRequest)
					return
				}
				libroExiste = true
				libros[i].Descargado = true
				break
			}
		}

		if !libroExiste {
			http.Error(w, "Libro no encontrado", http.StatusBadRequest)
			return
		}
	}

	mutex.Lock()
	defer mutex.Unlock()

	ultimoID++
	descarga.ID = ultimoID
	descarga.Fecha = time.Now()
	descargas = append(descargas, descarga)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(descarga)
}

func listarDescargas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(descargas)
}

// ======================
// MAIN
// ======================

func main() {

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			crearUsuario(w, r)
		} else if r.Method == http.MethodGet {
			listarUsuarios(w, r)
		} else {
			http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/libros", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			crearLibro(w, r)
		} else if r.Method == http.MethodGet {
			listarLibros(w, r)
		} else {
			http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/descargas", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			crearDescarga(w, r)
		} else if r.Method == http.MethodGet {
			listarDescargas(w, r)
		} else {
			http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
