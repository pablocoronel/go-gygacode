package main

import "fmt"

// estructura
type persona struct {
	nombre string
	edad   int
}

// funcion que calcula la persona mas vieja
func older(p1, p2 persona) (persona, int) {
	if p1.edad > p2.edad {
		return p1, (p1.edad - p2.edad)
	}

	return p2, (p2.edad - p1.edad)
}

func main() {
	// personas
	bob := persona{nombre: "Bob", edad: 27}
	patricio := persona{nombre: "Patricio", edad: 40}
	calamardo := persona{nombre: "Calamardo", edad: 45}

	// calculo de personas mas viejas
	bobPatricio, diferenciaBobPat := older(bob, patricio)
	bobCalamardo, diferenciaBobCala := older(bob, calamardo)
	patricioCalamardo, diferenciaPatCala := older(patricio, calamardo)

	// resultados
	fmt.Printf("Entre %s y %s es más viejo %s por %d años \n", bob.nombre, patricio.nombre, bobPatricio.nombre, diferenciaBobPat)

	fmt.Printf("Entre %s y %s es más viejo %s por %d años \n", bob.nombre, calamardo.nombre, bobCalamardo.nombre, diferenciaBobCala)

	fmt.Printf("Entre %s y %s es más viejo %s por %d años", patricio.nombre, calamardo.nombre, patricioCalamardo.nombre, diferenciaPatCala)
}
