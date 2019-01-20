package main

import "fmt"

// Nuestro type personal
type dinero int

// A un tipo personal se le pueden agregar metodos
// antes del nombre lleva un "recibidor"
// Sobreescribo la funcion String de int
func (d dinero) String() string {
	return fmt.Sprintf("$%d", d)
}

func main() {
	// una variable del tipo "dinero"
	var sueldo dinero
	sueldo = 25000
	fmt.Println("Sueldo es: ", sueldo)

	// a sueldo ya no se le puede sumar un INT, xq son tipos distintos ahora
	aumento := 30000                        //Es un int normal
	nuevoSueldo := sueldo + dinero(aumento) // convierto el int a el tipo "dinero"
	fmt.Println("El sueldo con aumeto es:", nuevoSueldo)
}
