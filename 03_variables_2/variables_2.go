package main

import "fmt"

func main() {
	// Nombrado de variables
	// Empieza con letra, o guion bajo, puede contener numeros
	numero := 25
	_nombre := "Pablo"
	apellido1 := "Coronel"

	// Debe ser camelCase
	nombreUsuario := "Admin"

	// Se diferencia entre mayuscula y minuscula, son dos variables distintas
	Numero := 40

	// debe tener identado si es la misma lineam, para continuar
	fmt.Println(numero, _nombre, apellido1,
		nombreUsuario, Numero)

	// Declaracion multiple (el tipo se infiere):
	var (
		persona      = "Carlos"
		edad1, edad2 = 20, 21
	)

	fmt.Println(persona, edad1, edad2)

	// sin := tambien infiere el tipo
	var direccion = "Callao"
	fmt.Println(direccion)

	// Scope de variables (alcance)

}
