package main

import (
	"fmt"
)

func main() {
	// Diccionario de códigos → descripción
	codigos := map[string]string{
		"10": "notebook",
		"15": "tv",
		"21": "heladera",
		"27": "monitor",
		"35": "camara",
	}

	var resultados []string
	var input string

	fmt.Println("Ingrese códigos (0 para terminar):")

	for {
		fmt.Scanln(&input)

		if input == "0" {
			break
		}

		// Buscar en el mapa
		if desc, existe := codigos[input]; existe {
			resultados = append(resultados, desc)
		} else {
			resultados = append(resultados, "No encontrado")
		}
	}

	fmt.Println("Los valores del array son:", resultados)
}
