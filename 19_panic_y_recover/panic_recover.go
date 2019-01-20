package main

import "fmt"

func imprimir() {
	fmt.Println("Hola desde imprimir")

	// Si hubiese un error...
	defer func() {
		error := recover()
		fmt.Println(error)
	}()
	panic("Hubo un error")
}

func main() {
	fmt.Println("Hola desde main()")
	imprimir()
	fmt.Println("Vuelvo a main")
}
