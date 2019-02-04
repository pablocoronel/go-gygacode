package main

import (
	"fmt"
	"strconv"
	"time"
)

func enviarMensaje(out chan<- string, numero int) {
	for {
		// envia el mensaje
		out <- "Mensaje: " + strconv.Itoa(numero) //itoa convierte de int a string
		fmt.Println("Envieando mensaje func:", numero)
	}
}

func imprimir(in <-chan string) {
	for x := range in {
		fmt.Println(x)              //mostrar el mensaje recibido
		time.Sleep(1 * time.Second) //frena un segundo
	}
}

func main() {

	ch := make(chan string, 5) // 5 es la capacidad del mensajes del channel

	for i := 1; i < 5; i++ {
		go enviarMensaje(ch, i) //envio la gorutina
	}

	// ver en consola
	imprimir(ch)

}
