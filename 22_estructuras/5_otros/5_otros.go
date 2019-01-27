package main

import "fmt"

// por defecto los valores de los campos es 0 (cero) o "" si es string
type punto3D struct {
	x, y, z int
}

// una estructura vacia se puede usar para agregarle metodos
type paraMetodos struct{}

// una estructura NO puede tener un campo de la misma estructura, si un puntero
// type otro struct {
// 	otro
// }

// con puntero a si misma
type otro struct {
	numero int
	*otro
}

// comparable
type punto struct {
	a int
	b int
}

func main() {
	p := punto3D{}

	fmt.Println("estructura vacia:", p)

	// llamando a un puntero de la misma estructura
	h := otro{
		5,
		&otro{
			7,
			nil, //nil es el por defecto de un puntero
		},
	}

	fmt.Println("ver una estructura con un puntero a la misma estructura", h)

	// comparando estructuras
	a := punto{5, 7}
	b := punto{8, 9}
	c := punto{5, 7}

	fmt.Println("Comparacion de estructuras A con B;", a == b)
	fmt.Println("Comparacion de estructuras A con C:", a == c)
}
