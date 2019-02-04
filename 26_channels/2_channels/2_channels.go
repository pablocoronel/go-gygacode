package main

import (
	"fmt"
	"time"
)

func main() {
	channelNumero := make(chan int)
	channelCuadrado := make(chan int)

	// genera un numero
	go func() {
		for i := 0; ; i++ {
			channelNumero <- i
		}
	}() //los parentesis finales sirven para ejecutar la funcion anonima apenas termina la declaracion

	// eleva al cuadrado el numero
	go func() {
		for {
			x := <-channelNumero     //guarda en X lo que recibe del channel
			channelCuadrado <- x * x // envia al channel cuadrado la multiplicacion
		}
	}() //los parentesis finales sirven para ejecutar la funcion anonima apenas termina la declaracion

	// imprime en main // main tambien es una gorutina
	for {
		fmt.Println(<-channelCuadrado) //imprime lo que recibe el channel
		time.Sleep(1 * time.Second)    //frene el bucle for por 1 segundo
	}
}
