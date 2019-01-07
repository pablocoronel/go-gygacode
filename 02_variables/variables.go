package main

import "fmt"

func main() {
	// Forma larga
	var numero int //Por defecto tiene valor 0 (cero)

	numero = 25

	fmt.Println(numero)

	numero = 40

	fmt.Println(numero)

	// Forma corta
	nombre := "Pablo"
	fmt.Println(nombre)

	// Asignar varias variables en una linea
	nombre, numero = "Andrea", 26
	fmt.Println(nombre, numero)

	fmt.Println("cambio de valor entre variables:")

	nombre1 := "Pablo"
	nombre2 := "Manuel"

	fmt.Println("Antes: ", nombre1, nombre2)

	nombre1, nombre2 = nombre2, nombre1
	fmt.Println("Despues:", nombre1, nombre2)

	// asignacion multiple con nueva y ya existente variable
	nombre3, nombre := "Maria", "Carlos"
	fmt.Println(nombre3, nombre)
}
