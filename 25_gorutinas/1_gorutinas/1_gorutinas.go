package main

import (
	"fmt"
	"math/rand" //numero aleatorio
	"sync"      //manejar gorutinas
	"time"      //tiempo
)

// indeicar al programa cuando terminan las gorutinas que se estan ejecutando
var wg sync.WaitGroup

func main() {
	// indica la cantidad de gorutinas que debe esperar a que terminen para que contunue el resto del programa
	wg.Add(2)
	fmt.Println("Empieza...")

	go imprimir("A") //go dice que ésta funcion se tiene que ejecutar como una gorutina
	go imprimir("B")

	fmt.Println("Esperando a que terminen las gorutinas")

	// aca se queda esperano la finalizacion de las gorutinas para contunuar
	wg.Wait() //Cuando wg llega a cero (0) continua el resto, pasa a la linea siguiente

	fmt.Println("Ya terminó")
}

// funcion que se va a usar como Gorutina
func imprimir(etiqueta string) {
	// le resta 1 a wg (avisa que ya terminó la ejecucion de ésta gorutina)
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		// para que se entienda la ejecucion, se ponen unos milisegundos de espera entre cada iteracion, para que se vea como se ejecutan concurrentemente ambas funciones indicadas como gorutinas

		espera := rand.Intn(1000) // devuelve un numero entero aleatorio entre 0 y el numero indicado

		// retraaso de tiempo en cada iteracion
		// * Sleep pausa la ejecucion el tiempo indicado
		// * Duration es un numero que Sleep espera como parametro
		// * Milisecond es una constante de unidad tiempo a convertir la espera
		time.Sleep(time.Duration(espera) * time.Millisecond)

		fmt.Println("Etiqueta Nro", i, " de ", etiqueta)
	}
}
