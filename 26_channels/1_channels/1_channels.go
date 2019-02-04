package main

import (
	"fmt"
	"time"
)

func main() {
	// crea el channel
	miChannel := make(chan string)

	// usamos las funciones como Gorutines
	go enviarMensaje(miChannel)  // recibe un channel para poder comunicarse con la otra gorutina
	go recibirMensaje(miChannel) // recibe un channel para poder comunicarse con la otra gorutina

	var entrada string
	fmt.Scanln(&entrada)
	fmt.Println("entrada para cortar las gorutinas:", entrada)
	fmt.Println("fin")

}

func enviarMensaje(c chan string) {
	for {
		c <- "un mensaje"
	}
}

func recibirMensaje(c chan string) {
	var contador int

	for {
		contador++
		fmt.Println("Ver lo recibido:", <-c, "número:", contador)
		time.Sleep(time.Second * 1)
	}
}
