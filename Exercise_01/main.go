package main

import "fmt"

func main() {
	license := true
	age := 19

	ok := false

	if age > 15 {
		if license == true {
			ok = true
		}
	}

	if ok {
		fmt.Println("Puede seguir avanzando")
	} else {
		fmt.Println("No puede seguir circulando")
	}
}
