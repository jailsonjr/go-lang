package main

// EXERCÍCIO
// Desafio surpresa!
// Format Printing:
// - Decimal       %d
// - Hexadecimal   %#x
// - Unicode       %#U
// - Tab           \t
// - Linha nova    \n
// Faca um loop dos numeros 33 a 122, e utilize format printing para
// demonstra-los como texto/string

import (
	"fmt"
)

func main() {

	for x := 33; x <= 122; x++ {
		fmt.Printf("Value: %v %c \n", x, x)
	}

}
