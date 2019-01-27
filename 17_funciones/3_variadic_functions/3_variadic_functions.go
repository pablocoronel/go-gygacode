package main

import "fmt"

// La funcion tiene indefinida cantidad de parametros (0, 1 o muchos)
// Se defina poniendo 3 puntos antes del tipo del parametro
func sumar(numeros ...int) (resultado int) {
	for _, value := range numeros {
		resultado += value
	}

	return
}

// Un parametro obligatorio y varios indefinidos
func mostrar(palabra1 string, muchasPalabras ...string) string {
	for _, value := range muchasPalabras {
		palabra1 += " " + value
	}

	return palabra1
}

func main() {
	// Admite cualquier cantidad de parametros
	fmt.Println(sumar(1, 2, 3))
	fmt.Println(sumar(1, 2))
	fmt.Println(sumar(1))
	fmt.Println(sumar())

	// Si se le pasa un slice, hay que abrirlo con el Spread Operator
	miSlice := []int{
		1, 2, 3, 4, 5,
	}
	fmt.Println("Con slice:")
	fmt.Println(sumar(miSlice...))

	// Con 1 parametro obligatorio y muchos adicionales
	fmt.Println("Con uno obligatorio y muchos adicionales")
	fmt.Println(mostrar("Hola", "mundo", "desde", "funcion mostrar()"))
}
