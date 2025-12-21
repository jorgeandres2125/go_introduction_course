package main

import (
	"fmt"
)

func main() {
	var valores []int // slice dinámico para almacenar los números
	var input int

	fmt.Println("Ingrese valores (0 para terminar):")

	for {
		fmt.Scanln(&input) // lee un número desde consola

		if input == 0 {
			break // si es cero, salimos del bucle
		}

		valores = append(valores, input) // agregamos al slice
	}

	fmt.Println("Los valores del array son:", valores)
}
