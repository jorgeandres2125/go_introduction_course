package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

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

	fmt.Println()

	const myFirstConst = "a12"
	fmt.Println("My const is: ", myFirstConst)

	const myIntConst = 12
	fmt.Println("My const is: ", myIntConst)

	const myStringConst string = "aaa"
	fmt.Println("My const is: ", myStringConst)

	fmt.Println()
	var my8BitIntVar int8
	fmt.Printf("Valor de variable es: %d\n ", my8BitIntVar)

	fmt.Printf("type :%T, byte: %d, bits: %d\n", my8BitIntVar, unsafe.Sizeof(my8BitIntVar), unsafe.Sizeof(my8BitIntVar)*8)

	var my16BitIntVar int16
	fmt.Printf("type :%T, byte: %d, bits: %d\n", my16BitIntVar, unsafe.Sizeof(my16BitIntVar), unsafe.Sizeof(my16BitIntVar)*8)

	var my32BitIntVar int32
	fmt.Printf("type :%T, byte: %d, bits: %d\n", my32BitIntVar, unsafe.Sizeof(my32BitIntVar), unsafe.Sizeof(my32BitIntVar)*8)

	var my64BitIntVar int64
	fmt.Printf("type :%T, byte: %d, bits: %d\n", my64BitIntVar, unsafe.Sizeof(my64BitIntVar), unsafe.Sizeof(my64BitIntVar)*8)

	var myFloat32Var float32
	fmt.Printf("type :%T, byte: %d, bits: %d\n", myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myFloat64Var float64
	fmt.Printf("type :%T, byte: %d, bits: %d\n", myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

	fmt.Printf("String default value %s", myStringVar)

	myStringVar5 := `My string variable in golang
    with backticks`
	fmt.Printf("The variable is %s\n", myStringVar5)

	{
		fmt.Println("This is a new block")
		floatVar := 33.11
		fmt.Printf("type %T, Value: %f\n", floatVar, floatVar)
		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("type %T, Value: %s\n", floatStrVar, floatStrVar)

		intVar := 22
		fmt.Printf("type %T, Value: %d\n", intVar, intVar)
		intStrVar := fmt.Sprintf("%.d", intVar)
		fmt.Printf("type %T, Value: %s\n", intStrVar, intStrVar)

		intVal1, err := strconv.ParseInt("1333", 0, 64)

		fmt.Println(err)
		fmt.Printf("type %T, Value: %d\n", intVal1, intVal1)

		intVal2, err := strconv.ParseInt("aa122", 0, 64)

		fmt.Println(err)
		fmt.Printf("type %T, Value: %d\n", intVal2, intVal2)

		floatVar2, _ := strconv.ParseFloat("-11.2", 64)

		fmt.Printf("type %T, Value: %f\n", floatVar2, floatVar2)
	}
}
