package main

import "fmt"

// estructura superior
type persona struct {
	nombre   string
	apellido string
}

// estructura "hija"
type estudiante struct {
	persona
	profesion string
}

// estrucura que tiene campo con el mismo nombre, se le da prioridad al campo del "Padre", para acceder al campo del hijo se llama a la estructura Hija y luego al campo
type profesor struct {
	estudiante
	profesion string
}

func main() {

	// declaracion
	pablo := estudiante{
		persona{
			nombre:   "Pablo",
			apellido: "Coronel"},
		"Programador",
	}

	// acceso a los atributos
	fmt.Println("Estructura de Pablo:", pablo)
	fmt.Println("acceso a campos hijos: Nombre:", pablo.nombre)
	fmt.Println("acceso a campos hijos: Apellido:", pablo.apellido)
	fmt.Println("campo padre: Profesion:", pablo.profesion)

	laura := profesor{
		estudiante{
			persona{
				nombre:   "Laura",
				apellido: "Gonzalez",
			},
			"Programadora",
		},
		"Informatica",
	}

	fmt.Println("************** campos padre ***************")

	// acceso a campos, el padre tiene prioridad
	fmt.Println("nombre:", laura.nombre)
	fmt.Println("campo de mismo nombre (profesion)", laura.profesion)

	fmt.Println("************* campos hijo *****************")
	fmt.Println("campo hijo (profesion)", laura.estudiante.profesion)

	// declaracion de estructuras en una linea
	jose := profesor{estudiante{persona{nombre: "jose", apellido: "garcia"}, "soldador"}, "metalurgico"}
	fmt.Println("estructura en una linea", jose)

	// declaracion y luego asignacion
	var pedro profesor
	pedro.nombre = "Pedro"
	pedro.apellido = "Martinez"
	pedro.profesion = "vendedor"
	pedro.estudiante.profesion = "comercio"
	fmt.Println(pedro)
}
