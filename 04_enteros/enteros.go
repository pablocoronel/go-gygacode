package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Enteros con signo (negativos y positivos)
	// var entero8 int8
	// var entero16 int16
	// var entero32 int32
	// var entero64 int64

	// // enteros sin signo (solamente positivos)
	// var uentero8 uint8
	// var uentero16 uint16
	// var uentero32 uint32
	// var uentero64 uint64

	// // Alias
	// var enteroByte byte //equivale a uint8
	// var enteroRune rune //equivale a int32

	// // dependen de la plataforma
	// var enteroUint uint       //32 o 64 bits
	// var enteroInt int         //32 o 64 bits
	// var enteroUnitPtr uintptr //para valor de punteros

	// Conversiones de tipos
	// var edad int8 = 16
	// nuevaEdad := int32(edad)

	// tamaño en bytes de variables 1 byte = 8 bits
	var edad2 int32 = 25
	fmt.Println("Bytes: ", unsafe.Sizeof(edad2)) // 4 bytes =  4 * 8 = 32 bits
}
