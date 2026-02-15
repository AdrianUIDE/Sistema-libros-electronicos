/*
@autor: Adrian Herrera
@fecha: 15/02/2026
@version: 1.0
@descripcion: Sistema de Gestion de Libros Electronicos
*/

package main

import (
	"errors"
	"fmt"
	"time"
)

// =======================
// CREAREMOS LA CLASE USUARIO
// =======================
type Usuario struct {
	id        int
	nombres   string
	apellidos string
	email     string
}

func (u *Usuario) GetId() int {
	return u.id
}

func (u *Usuario) SetId(id int) {
	u.id = id
}

func (u *Usuario) GetNombres() string {
	return u.nombres
}

func (u *Usuario) SetNombres(nombres string) {
	u.nombres = nombres
}

func (u *Usuario) GetApellidos() string {
	return u.apellidos
}

func (u *Usuario) SetApellidos(apellidos string) {
	u.apellidos = apellidos
}

func (u *Usuario) GetEmail() string {
	return u.email
}

func (u *Usuario) SetEmail(email string) {
	u.email = email
}

// =======================
// CREAREMOS LA CLASE LIBRO ELECTRONICO
// =======================
type Libro struct {
	id         int
	titulo     string
	autor      string
	isbn       string
	descargado bool
}

func (l *Libro) GetId() int {
	return l.id
}

func (l *Libro) SetId(id int) {
	l.id = id
}

func (l *Libro) GetTitulo() string {
	return l.titulo
}

func (l *Libro) SetTitulo(titulo string) {
	l.titulo = titulo
}

func (l *Libro) GetAutor() string {
	return l.autor
}

func (l *Libro) SetAutor(autor string) {
	l.autor = autor
}

func (l *Libro) GetIsbn() string {
	return l.isbn
}

func (l *Libro) SetIsbn(isbn string) {
	l.isbn = isbn
}

func (l *Libro) GetDescargado() bool {
	return l.descargado
}

func (l *Libro) SetDescargado(descargado bool) {
	l.descargado = descargado
}

// =======================
// CREAREMOS LA CLASE DESCARGA
// =======================
type Descarga struct {
	id      int
	usuario *Usuario
	libros  []*Libro
	fecha   time.Time
}

func (d *Descarga) GetId() int {
	return d.id
}

func (d *Descarga) SetId(id int) {
	d.id = id
}

func (d *Descarga) GetUsuario() *Usuario {
	return d.usuario
}

func (d *Descarga) SetUsuario(usuario *Usuario) {
	d.usuario = usuario
}

func (d *Descarga) GetLibros() []*Libro {
	return d.libros
}

func (d *Descarga) SetLibros(libros []*Libro) {
	d.libros = libros
}

func (d *Descarga) GetFecha() time.Time {
	return d.fecha
}

func (d *Descarga) SetFecha(fecha time.Time) {
	d.fecha = fecha
}

// =======================
// CREAREMOS LA CLASE SISTEMA DE LIBROS
// =======================
type SistemaLibros struct {
	usuarios  []*Usuario
	libros    []*Libro
	descargas []*Descarga
	lastId    int
}

// Registrar usuario
func (s *SistemaLibros) AgregarUsuario(nombres, apellidos, email string) error {
	if nombres == "" || apellidos == "" || email == "" {
		return errors.New("Datos invalidos")
	}

	s.lastId++
	usuario := &Usuario{}
	usuario.SetId(s.lastId)
	usuario.SetNombres(nombres)
	usuario.SetApellidos(apellidos)
	usuario.SetEmail(email)

	s.usuarios = append(s.usuarios, usuario)
	fmt.Println("Usuario registrado correctamente")
	return nil
}

// Registrar libro
func (s *SistemaLibros) AgregarLibro(titulo, autor, isbn string) error {
	if titulo == "" || autor == "" || isbn == "" {
		return errors.New("Datos invalidos")
	}

	s.lastId++
	libro := &Libro{}
	libro.SetId(s.lastId)
	libro.SetTitulo(titulo)
	libro.SetAutor(autor)
	libro.SetIsbn(isbn)
	libro.SetDescargado(false)

	s.libros = append(s.libros, libro)
	fmt.Println("Libro electronico registrado correctamente")
	return nil
}

// Descargar libro
func (s *SistemaLibros) DescargarLibro(usuarioId int, librosIds ...int) error {
	var usuario *Usuario
	var libros []*Libro

	for _, u := range s.usuarios {
		if u.GetId() == usuarioId {
			usuario = u
			break
		}
	}

	if usuario == nil {
		return errors.New("Usuario no encontrado")
	}

	for _, idLibro := range librosIds {
		var libro *Libro
		for _, l := range s.libros {
			if l.GetId() == idLibro {
				libro = l
				break
			}
		}

		if libro == nil {
			return errors.New("Libro no encontrado")
		}

		if libro.GetDescargado() {
			return errors.New("El libro ya fue descargado")
		}

		libros = append(libros, libro)
	}

	s.lastId++
	descarga := &Descarga{}
	descarga.SetId(s.lastId)
	descarga.SetUsuario(usuario)
	descarga.SetLibros(libros)
	descarga.SetFecha(time.Now())

	for _, libro := range libros {
		libro.SetDescargado(true)
	}

	s.descargas = append(s.descargas, descarga)
	fmt.Println("Descarga realizada correctamente")
	return nil
}

// ===========================
// Funcion MENU y PRINCIPAL
// ===========================
func menu() {
	fmt.Println("===== Sistema de Libros Electronicos =====")
	fmt.Println("1. Registrar usuario")
	fmt.Println("2. Registrar libro")
	fmt.Println("3. Descargar libro")
	fmt.Println("4. Listar usuarios")
	fmt.Println("5. Listar libros")
	fmt.Println("6. Listar descargas")
	fmt.Println("7. Salir")
	fmt.Print("Seleccione una opcion: ")
}

func main() {
	sistema := &SistemaLibros{}

	for {
		menu()
		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			fmt.Println("Agregar nuevo usuario: ")
			var n, a, e string
			fmt.Print("Nombres: ")
			fmt.Scan(&n)
			fmt.Print("Apellidos: ")
			fmt.Scan(&a)
			fmt.Print("Email: ")
			fmt.Scan(&e)
			sistema.AgregarUsuario(n, a, e)

		case 2:
			fmt.Println("Agregar nuevo libro: ")
			var t, au, isbn string
			fmt.Print("Titulo: ")
			fmt.Scan(&t)
			fmt.Print("Autor: ")
			fmt.Scan(&au)
			fmt.Print("ISBN: ")
			fmt.Scan(&isbn)
			sistema.AgregarLibro(t, au, isbn)

		case 3:
			fmt.Println("Descargar libro")
			var usuarioId, libroId int
			fmt.Print("Ingrese el ID del usuario: ")
			fmt.Scan(&usuarioId)
			fmt.Print("Ingrese el ID del libro: ")
			fmt.Scan(&libroId)
			if err := sistema.DescargarLibro(usuarioId, libroId); err != nil {
				fmt.Println("Error:", err)
			}
		case 4:
			fmt.Println("Listado de usuarios")
			for _, usuario := range sistema.usuarios {
				fmt.Println(
					"ID:", usuario.GetId(),
					"Nombres:", usuario.GetNombres(),
					"Apellidos:", usuario.GetApellidos(),
					"Email:", usuario.GetEmail(),
				)
			}
		case 5:
			fmt.Println("Listado de libros electronicos")
			for _, libro := range sistema.libros {
				fmt.Println(
					"ID:", libro.GetId(),
					"Titulo:", libro.GetTitulo(),
					"Autor:", libro.GetAutor(),
					"ISBN:", libro.GetIsbn(),
					"Descargado:", libro.GetDescargado(),
				)
			}
		case 6:
			fmt.Println("Listado de descargas")
			for _, descarga := range sistema.descargas {
				fmt.Println(
					"ID:", descarga.GetId(),
					"Usuario:", descarga.GetUsuario().GetNombres(),
					"Fecha:", descarga.GetFecha(),
				)
				for _, libro := range descarga.GetLibros() {
					fmt.Println("   Libro:", libro.GetTitulo())
				}
			}
		case 7:
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opcion no valida")
		}
	}
}
