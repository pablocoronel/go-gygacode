package main

import "fmt"

func main() {

	var (
		encabezado = `
		**********************************
		CONTADOR DE NúMEROS IMPARES
		**********************************
		`

		textoResultado = `
		**********************************
		RESULTADO:
		*********************************
		`
	)

	var (
		numero1, numero2, contadorImpares int
	)

	fmt.Println(encabezado)

	fmt.Println("Ingrese el primer número:")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese el segundo número:")
	fmt.Scanln(&numero2)

	fmt.Println("**************************")
	// Lógica:
	for i := numero1; i <= numero2; i++ {

		if i%2 != 0 {
			contadorImpares++
			fmt.Println(i, "es un número impar")
		}
	}

	fmt.Println(textoResultado)
	fmt.Printf("Entre el número %d y el número %d hay: %d números impares", numero1, numero2, contadorImpares)
}
