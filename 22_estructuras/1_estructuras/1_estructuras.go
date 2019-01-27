package main

import "fmt"

// Persona: estructura
type humano struct {
	Nombre string
	Edad   int
}

func main() {
	// 1)____ declarar variable del tipo humano
	var p humano

	p.Nombre = "Pablo"
	p.Edad = 27

	fmt.Println("estructura del tipo humano: ", p)
	fmt.Println("Nombre p:", p.Nombre)
	fmt.Println("Edad p: ", p.Edad)

	// 2)_____ declarar variable del tipo humano
	p2 := humano{Nombre: "Cacho", Edad: 57}
	fmt.Println("Variable 2: ", p2)
	fmt.Println("Variable 2, nombre: ", p2.Nombre)
	fmt.Println("Variable 2, edad", p2.Edad)

	// 3)_____ asignando directamente los valores
	p3 := humano{"Roberto", 79}
	fmt.Println("3ra forma: ", p3)
	fmt.Println("3ra forma nombre:", p3.Nombre)
	fmt.Println("3ra forma edad:", p3.Edad)
}
