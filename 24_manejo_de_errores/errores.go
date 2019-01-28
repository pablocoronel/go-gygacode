package main

import (
	"errors"
	"fmt"
)

// puede ser adentro de var() o individuales
var (
	// ErrorUsuarioNoValido ... //el comentario lo pide un plugin
	ErrorUsuarioNoValido = errors.New("El usuario no es valido")

	// ErrorOtro ...
	ErrorOtro = errors.New("Otro error")
)

// devolviendo un error
func baneado(usuario string) (err error) {
	ban := false

	switch usuario {
	case "miguel":
		ban = true
	case "carlos":
		ban = false

	case "juan":
		return ErrorUsuarioNoValido //usando los errores personalizados

	case "pedro":
		return fmt.Errorf("Usuario en proceso de registro") // Errorf devuelve un error
	default:
		return fmt.Errorf("algo paso")
	}

	if !ban {
		fmt.Printf("El usuario %s no está baneado \n", usuario)
	} else {
		fmt.Printf("El usuario %s está baneado \n", usuario)
	}

	return nil
}

// funcion para manejar los errores
func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	// Modo recomendado por GO
	err := baneado("miguel")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = baneado("carlos")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = baneado("juan")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = baneado("pedro")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = baneado("que se yo")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Modo recomendado por la comunidad
	fmt.Println("******* modo con funcion que maneja los errores **************")

	err = baneado("miguel")
	checkError(err)

	err = baneado("carlos")
	checkError(err)

	err = baneado("juan")
	checkError(err)

	err = baneado("pedro")
	checkError(err)

	err = baneado("que se yo")
	checkError(err)
}
