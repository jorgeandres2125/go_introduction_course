package main

import (
	"fmt"
)

func main() {
	ml := make(map[int]string)
	ml[1] = "A"
	ml[2] = "B"
	ml[3] = "C"

	fmt.Println(ml)
	fmt.Println(ml[1])

	delete(ml, 2)
	fmt.Println(ml)

	ml[1] = "A2"
	fmt.Println(ml)

	ml[7] = ""
	fmt.Println(ml)
	fmt.Println(ml[7])
	fmt.Println(ml[99])

	v, ok := ml[7]
	fmt.Println("The value:", v, "Present?", ok)

	v, ok = ml[99]
	fmt.Println("The value:", v, "Present?", ok)

	m2 := map[int]string{
		4: "A",
		5: "C",
		7: "Z",
	}
	fmt.Println(m2)
}
