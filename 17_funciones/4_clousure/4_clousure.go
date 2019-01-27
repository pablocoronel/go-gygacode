package main

import "fmt"

// Funcion normal
func imprimir(texto string) {
	fmt.Println("Desde Imprimir()", texto)
}

// Una funcion puede recibir como parametro otra funcion
func imprimir2(otraFuncion func(string)) {
	otraFuncion("Hola")
}

// Funcion clousure (va a devolver otra funcion (que a su vez devuelve un int))
func incrementar() func() int {
	i := 0

	return func() (r int) {
		r = i
		i += 2
		return
	}
}

//
func incremento() {
	var i int
	i++
	fmt.Println(i)
}

func main() {
	nombre := "Pablo Manuel Coronel"
	// 1)____ Asigno localmente en una variable la funcion externa)
	// Si dos funciones tienen el mismo tipo (la misma firma) mas adelante se puede asignar en la misma variable
	imprimirEnVariable := imprimir
	imprimirEnVariable("pablo")

	// 2)____ Declaro otra funcion dentro de una funcion
	otraFuncion := func() {
		fmt.Println("Desde adentro de una funcion:", nombre)
	}
	otraFuncion()

	// 3)_____ uso una funcion que tiene otra funcion como parametro
	imprimir2(imprimir)

	// 4)_____ llamado a clousure (se mantienen los valores de las variables de la funcion externa)
	inc := incrementar()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	// Sin guardar el estado interno
	fmt.Println("Sin guardar el estado interno:")
	incremento()
	incremento()
	incremento()
}
