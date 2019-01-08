package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Solo se declaran strings con comillas dobles (o acento invertido, go lo toma literal (tal cual))
	// Un string es un array de caracteres
	var cadena string

	cadena = "Hola mundo"

	fmt.Println(cadena)

	// tamaño
	fmt.Println(len(cadena))

	fmt.Println(cadena[3]) // La salida es en unicode

	// Tomar un segmento de string
	fmt.Println(cadena[0:4]) //cada posision del : es opcional, por defecto el primero y el ultimo respectivamente

	// Concatenar
	cadena = cadena + " nuevamente"
	fmt.Println(cadena)

	nombre := "Pablo"

	// asi tambien se concatena de forma abreviada
	nombre += " Coronel"
	fmt.Println(nombre)

	literal := `<html>`
	fmt.Println(literal)

	// escapar caracteres
	textoConComillas := "Hola \"entre commilas\""
	fmt.Println(textoConComillas)

	// Convertir a string
	fmt.Println(textoConComillas + " " + strconv.Itoa(25))
}
