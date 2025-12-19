package main

import (
	"fmt"
)

func main() {
	yearsOld := 32

	fmt.Println("Operators")
	fmt.Println(yearsOld > 30)  // true
	fmt.Println(yearsOld < 32)  // false
	fmt.Println(yearsOld <= 32) // true
	fmt.Println(yearsOld >= 40) // false
	fmt.Println(yearsOld == 32) // true

	fmt.Println()
	// OR

	fmt.Println(yearsOld < 32 || yearsOld == 32) // (false or true) = true
	fmt.Println(yearsOld < 32 || yearsOld == 33) // (false or false) = false
	fmt.Println(yearsOld < 40 || yearsOld == 33) // (true or true) = true

	fmt.Println()
	// AND

	fmt.Println(yearsOld < 32 && yearsOld == 32) // (false and true) = false
	fmt.Println(yearsOld < 32 && yearsOld == 33) // (false and false) = false
	fmt.Println(yearsOld < 40 && yearsOld == 32) // (true and true) = true

	fmt.Println()

	// NOT

	fmt.Println(true)             // true
	fmt.Println(!true)            // false
	fmt.Println(yearsOld < 40)    // true
	fmt.Println(!(yearsOld < 40)) // false

}
