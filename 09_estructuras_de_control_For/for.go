package main

import "fmt"

func main() {
	// 1) For simil while

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// 2) for normal
	for j := 0; j < 4; j++ {
		fmt.Println(j)
	}
}
