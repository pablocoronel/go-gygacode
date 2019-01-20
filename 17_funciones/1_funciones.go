package main

import "fmt"

// Funcion con parametro
func imprimeNombre(nombre string) {
	fmt.Println("Afuera de main")
	fmt.Println("nombre: ", nombre)
}

// Funcion con return 1
func suma(numero1 int, numero2 int) int {
	return numero1 + numero2
}

// Funcion con return 2
func resta(numero1, numero2 int) (resultado int) {
	resultado = numero1 - numero2
	return
}

func main() {
	imprimeNombre("pablo")

	resultadoSuma := suma(10, 15)
	fmt.Println(resultadoSuma)
	resultadoResta := resta(55, 25)
	fmt.Println(resultadoResta)
}
