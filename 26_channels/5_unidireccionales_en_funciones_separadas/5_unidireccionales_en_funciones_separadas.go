package main

import (
	"fmt"
	"time"
)

// genera un numero
func generarNumero(out chan<- int) { // recibe un channel DE SALIDA,  (a la derecha)
	for i := 0; i <= 5; i++ {
		out <- i
	}

	close(out)
}

// recibe dos channels, el de ingreso (izquierda) y el de envio (derecha)
func elevarAlCuadrado(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}

	close(out)
}

func imprimir(in <-chan int) {
	for x := range in {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	channelNumero := make(chan int)
	channelCuadrado := make(chan int)

	// generar el numero
	go generarNumero(channelNumero)

	// eleva al cuadrado el numero
	go elevarAlCuadrado(channelNumero, channelCuadrado)

	// imprime en main
	imprimir(channelCuadrado)
}
