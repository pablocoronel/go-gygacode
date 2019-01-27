package main

import (
	"fmt"
	"math"
)

// estructura
type rectangulo struct {
	alto, ancho float64
}

type circulo struct {
	radio float64
}

// metodo para la estructura
// lo primero es un "recibidor" que indica a cual estructura pertenece y permite usar el valor recibido
func (r rectangulo) area() float64 {
	return r.alto * r.ancho
}

// a pesar de tener el mismo nombre, no hay conflicto xq petenece al ambito de la estructura asociada
func (c circulo) area() float64 {
	return c.radio * math.Pi
}

// metodo que devuelve una estructura nueva
func (r rectangulo) incrementar(i int) rectangulo {
	return rectangulo{
		alto:  r.alto * float64(i),
		ancho: r.ancho * float64(i),
	}
}

// metodo que opera sobre la misma estructura (usando puntero)
func (r *rectangulo) incrementarDirectamente(i int) {
	r.alto = r.alto * float64(i)
	r.ancho *= float64(i)
}

func main() {
	// declarar estructura
	r1 := rectangulo{10.0, 12.0}
	fmt.Println("Metodo de la estructura, area:", r1.area())

	c1 := circulo{radio: 10.0}
	fmt.Println("Area de un circulo:", c1.area())

	// usa la funcion que devuelve una nueva estructura
	rincrementado := r1.incrementar(5)
	fmt.Println("tamaño normal:", r1)
	fmt.Println("tamaño incrementado:", rincrementado)

	// usando puntero para modificar la misma estructura
	fmt.Println("Antes: ", r1)
	r1.incrementarDirectamente(2)
	fmt.Println("Despues", r1)
}
