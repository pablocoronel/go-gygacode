package main

import "fmt"

func main() {
	// Slice de strings
	nombres := []string{
		"Pablo",
		"Manuel",
	}

	// for con range (simil foreach)
	for index, value := range nombres {
		fmt.Printf("Posicion: %d, Nombre: %s \n", index, value)
	}

	// obviar el indice
	fmt.Println("****************")
	for _, value := range nombres {
		fmt.Printf("Nombre: %s \n", value)
	}

	// obviar el valor, solo obtener el indeice
	fmt.Println("****************")
	for index := range nombres {
		fmt.Println("indice de los nombres: ", index)
	}
}
