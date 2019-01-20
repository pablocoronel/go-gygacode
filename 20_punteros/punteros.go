package main

import "fmt"

// Funcion que recibe un puntero
func incrementar(numero *int) {
	*numero++
	fmt.Println("Usando el puntero =>", *numero)
}

func main() {
	nombre := "Pablo"
	// Un puntero es UN TIPO DE VALOR que contiene la direccion de momoria de una variable
	miPuntero := &nombre

	fmt.Println("El valor de nombre =>", nombre)

	// Con & se obtiene la direccion de memoria
	fmt.Println("El valor de la direccion de memoria de nombr =>", miPuntero)

	fmt.Println("El valor de la variable a la que apunta el puntero =>", *miPuntero)

	// Usando asterisco *
	*miPuntero = "Jorge"
	fmt.Println("Nuevo valor de la variable =>", nombre)

	nombre = "Claudia"
	fmt.Println("Segundo nuevo valor =>", *miPuntero)

	fmt.Println("********************************")
	unNumero := 1
	fmt.Println("Antes de usar el puntero =>", unNumero)
	// Con el & se le pasa la direccion de memoria de una variable, es decir, un puntero
	incrementar(&unNumero)
}
