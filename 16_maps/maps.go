package main

import "fmt"

func main() {
	// Los maps son simil diccionarios, tiene asociacion clave, valor
	// opcional, 3er parametro: capacidad
	a := make(map[string]string)

	a["nombre"] = "Pablo"

	fmt.Println("Ver map: ", a)

	// acceder al valor de un map
	fmt.Println("Ver valor de un elemento de map:", a["nombre"])

	// declarar map con valores:
	edades := map[string]int{
		"Pablo": 27,
		"Ana":   25,
	}

	fmt.Println("Ver mapa con valores:", edades)

	// Borrar elemento del map
	delete(edades, "Pablo")
	fmt.Println("Ver map con elemento borrado:", edades)

	// Los maps devuelven dos valores: 1ro: el valor, 2do: bool de existencia del elemento en el map
	edad, existe := edades["Lupe"]
	fmt.Printf("Valoes devueltos por elemento de map: valor => %d, existe => %t", edad, existe)
}
