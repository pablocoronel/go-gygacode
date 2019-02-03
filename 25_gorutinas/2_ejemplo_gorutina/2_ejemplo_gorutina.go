package main

import (
	"fmt"
	"time"
)

func main() {
	go animacion(100 * time.Millisecond) //es una gorutina sin WaitGroup, asi que se ejectuta y la siguiente linea tambien se sigue ejecutando, cuando termina main, ahi se termina la gorutina animacion

	// aparte de la gorutina, el programa sigue
	const numero = 45
	resultado := fibonacci(numero)
	fmt.Printf("\r Fibonacci(%d) = %d\n", numero, resultado) //despues de imprimir, se termina main, por lo tanto se termina Animacion
}

func animacion(retraso time.Duration) {
	// es un for infinito
	for {
		for _, value := range `\|/-` {
			fmt.Printf("\r%c", value)

			time.Sleep(retraso)
		}
	}
}

// serie de fibonacci, es una funcion recursiva, porque se llama a si misma
func fibonacci(numero int) int {
	if numero < 2 {
		return numero
	}

	return fibonacci(numero-1) + fibonacci(numero-2)
}
