package main

import "fmt"

func main() {
	// declarar un array:
	var x [3]int //x es un array de 3 posiciones que admite Integers
	fmt.Println("Array:")
	fmt.Println(x)

	// matriz
	var y [3][2]int //y es una matriz de 3x2
	fmt.Println("Matriz:")
	fmt.Println(y)

	// asignar un valor a una posicion de un array:
	x[1] = 25
	fmt.Println("Array con valor asignado:")
	fmt.Println(x[1])

	// acceder a una posicion determinada
	fmt.Println("Ver posicion en particular")
	fmt.Println(x[1])

	// declarar e inicializar
	z := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Ver array inicializado")
	fmt.Println(z)

	// El tamaño del array es inmutable, se tiene que declara de antemano
	t := [...]int{100, 200}
	fmt.Println("Array con tamaño basado en valores")
	fmt.Println(t)

	// declaracion de array verticalmente (facil comentado por coma al final)
	s := [...]int{
		1000,
		2000,
		3000,
	}
	fmt.Println("Ver declaracion vertical con , final")
	fmt.Println(s)

	// Tamaño de un array se ve con la funcion len
	fmt.Println("Tamaño de un array")
	fmt.Println(len(s))

	// Ultimo elemento de un array
	fmt.Println(s[len(s)-1])

}
