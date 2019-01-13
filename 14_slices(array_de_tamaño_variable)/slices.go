package main

import "fmt"

func main() {
	//Sslice es una porcion de un array, lo cual permite que pueda variar su tamaño, a diferencia de los arrays

	// Slice
	var j []int
	fmt.Println("Slice", j)

	// Slice con inicializacion
	k := []int{1, 2}
	fmt.Println("Slice con inicializacion", k)

	// Crear slice declarando tamaño
	var m = make([]int, 3)
	n := make([]int, 5, 10) //con 3er parametro se indica la Capacidad
	fmt.Println("Slice con make declarando el tamaño", m, n)

	// array
	miArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Array: ", miArray)

	// slice que apunta al array:
	var miSlice []int
	miSlice = miArray[2:5] //desde la posicion 2 al 5 sin incluirlo

	fmt.Println("Ver slice que apunta a array:", miSlice)

	// Cambiar un valor en el slice cambia el array, porque el slice APUNTA al array
	// NO crea un array nuevo
	miSlice[0] = 100
	fmt.Println("ver array modificado desde el slice:", miArray)

	// Copiar slices
	sliceDos := []int{0, 1}
	// el slice de receptor recibe los elementos copiados segun su tamaño
	// si tiene tamaño 3 y reciebe 7 elementos, copia solo los primeros 3
	// si tiene tamaño 3 y recibe 1, copia uno y mantiene los dos restantes "merge"
	copy(sliceDos, miSlice)
	fmt.Println("Ver copia de slices:", sliceDos)
}
