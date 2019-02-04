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
		for i := 0; i <= 5; i++ {
			channelNumero <- i
		}

		close(channelNumero)
	}()

	// eleva al cuadrado el numero
	go func() {
		// con for range la iteracion se detiene cuando ya no hay emision del canal
		for x := range channelNumero {
			channelCuadrado <- x * x
		}

		close(channelCuadrado)
	}()

	// imprime en main
	// con for range la iteracion se detiene cuando ya no hay emision del canal
	for x := range channelCuadrado {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}
