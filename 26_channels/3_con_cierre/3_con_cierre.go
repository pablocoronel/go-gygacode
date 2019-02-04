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

		close(channelNumero) //cierra el channel
	}()

	// eleva al cuadrado el numero
	go func() {
		for {
			x, estadoChannel := <-channelNumero //un channel devuelve un segundo valor, indica si el canal esta abierto (true) o cerrado (false)

			// si ya no esta transmitiendo el canal (esta cerrado), sale del for
			if !estadoChannel {
				break
			}

			channelCuadrado <- x * x
		}

		close(channelCuadrado) //tambien cerramos el segundo channel, xq el primero ya no emite
	}()

	// imprime en main // main tambien es una gorutina
	for {
		x, ok := <-channelCuadrado

		// aca tambien consultamos si dejo de emitir el canal
		if !ok {
			break
		}

		fmt.Println(x)              //imprime lo que recibe el channel
		time.Sleep(1 * time.Second) //frene el bucle for por 1 segundo
	}
}
