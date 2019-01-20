package main

import (
	"fmt"
	"os"
)

func primera() {
	fmt.Println("Primera")
}

func segunda() {
	fmt.Println("Segunda")
}

func tercera() {
	fmt.Println("Tercera")
}

func main() {
	// La primera funcion la ejecuta al final de todo el ambito de la funcion main()
	// defer primera()
	primera()
	segunda()
	tercera()

	// Abrir el archivo
	archivo, error := os.Open("texto.txt")

	if error != nil {
		panic(error)
	}
	// Lo coloco acá para que siempre se cierre el archivo al final
	defer archivo.Close()

	// Slice para almacenar lo que contiene el archivo de texto
	data := make([]byte, 413)
	// "vuelca" en data el contenido del archivo
	texto, error := archivo.Read(data)

	if error != nil {
		panic(error)
	}

	fmt.Printf("Cantidad de bytes leido: %d", texto)
	fmt.Printf("Texto: %q", data)
	fmt.Printf("Errores: %v", error)
	fmt.Println()

	// En un for, defer va guardando la ejecucion con los valores de ese momento
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
