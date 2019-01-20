package main

import "fmt"

// Funcion que retorna multiples valores
// Con retornos nombrados
func multiplicar(n1 int) (numero1, numero2, numero3 int) {
	numero1 = n1 * 10
	numero2 = n1 * 20
	numero3 = n1 * 30
	return
}

// Sin nombrar los retornos
func multitplicarDos(num int) (int, int, int) {
	numero1 := num * 10
	numero2 := num * 20
	numero3 := num * 30

	return numero1, numero2, numero3
}

func main() {
	// usando la funcion en PrintLn
	fmt.Println(multiplicar(10))

	// Asignando los valores devueltos a variables
	c1, c2, c3 := multiplicar(25)
	fmt.Println("Con valores de retorno guardados en variables")
	fmt.Println(c1, c2, c3)

	// Sin nombrar los valores de retorno
	fmt.Println("Sin valores de retorno nombrado")
	fmt.Println(multitplicarDos(15))

	// Ignorando variables que no voy a usar
	_, d2, _ := multiplicar(20)
	fmt.Println("Ignorando algunas de las variables devueltas")
	fmt.Println(d2)
}
