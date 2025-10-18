package main

// EXERCÍCIO 01
// Escreva um programa que mostre um numero decimal, binario e hexadecimal

import (
	"fmt"
)

func main() {
	x := 31337
	fmt.Printf("Value: %v", x)
	fmt.Println("")
	fmt.Println("")
	fmt.Printf("Binario: %d\n", x)
	fmt.Printf("Decimal: %#b\n", x)
	fmt.Printf("Hexadecimal: %#x\n", x)
}
