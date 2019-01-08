package main

import "fmt"

func main() {
	// if normal
	a := 5
	if a <= 6 {
		fmt.Println("a es menor o igual que 6")
	}

	// con sentencias adentro
	b := 5
	if b = b + 1; b < 7 {
		fmt.Println("B sigue siendo menor que 7")
	}

	// if, else if, else
	if a < b {
		fmt.Println("A es menor que B")
	} else if a > b {
		fmt.Println("A es mayor que B")
	} else {
		fmt.Println("A es igual que B")
	}
}
