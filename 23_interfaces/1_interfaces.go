package main

import "fmt"

type persona struct {
	nombre string
	email  string
	edad   int
}

type moderador struct {
	persona
	foro string
}

func (m moderador) abrirForo() {
	fmt.Println("abrir foro")
}

func (m moderador) cerrarForo() {
	fmt.Println("cerrar foro")
}

type administrador struct {
	persona
	seccion string
}

func (a administrador) crearForo() {
	fmt.Println("crear foro")
}

func (a administrador) eliminarForo() {
	fmt.Println("eliminar foro")
}

// ***********************************************
// funcion que usa dos funciones
func presentarse(p persona) {
	fmt.Println("Nombre:", p.getNombre())
	fmt.Println("Email:", p.getEmail())
}

func (p persona) getNombre() string {
	return p.nombre
}

func (p persona) getEmail() string {
	return p.email
}

// **********************************************
// interface
type usuario interface {
	getNombre() string
	getEmail() string
}

func presentarseConInterface(u usuario) {
	fmt.Println("Nombre:", u.getNombre())
	fmt.Println("Email:", u.getEmail())
}

func main() {
	pablo := persona{"Pablo", "pablo@gmail.com", 29}
	presentarse(pablo)

	juan := administrador{persona{"Juan", "juean@gmail.com", 25}, "deportes"}
	// da error xq presentarse solo admite el tipo persona, no administrador
	// para arrreglar eso se usan interfaces
	// presentarse(juan)

	// usando interface
	presentarseConInterface(juan)

	// variable del tipo de una interface
	var i usuario
	i = pablo //struct persona
	fmt.Println("variable del tipo de la interface:", i)
	fmt.Println("i nombre (persona):", i.getNombre())

	i = juan
	fmt.Println("variable del tipo administrador (usa persona):", i)
	fmt.Println("i nombre (administrador)", i.getNombre())
}
