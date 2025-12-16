package main

import "fmt"

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println("Valor de variable es: ", myIntVar)
	var myUintVar uint
	myUintVar = 12
	fmt.Println("Valor de variable es: ", myUintVar)

	var myStringVar string
	myStringVar = "Hola Mundo"
	fmt.Println("Valor de variable es: ", myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("Valor de variable es: ", myBoolVar)

	fmt.Println("Puntero ", &myStringVar)

	var myFloatVar float32
	myFloatVar = 3.14
	fmt.Println("Valor de variable es: ", myFloatVar)

	myIntVar2 := 12
	fmt.Println("Valor de variable es: ", myIntVar2)

	myStringVar2 := "My string variable with :="
	fmt.Println("Valor de variable es: ", myStringVar2)

}
