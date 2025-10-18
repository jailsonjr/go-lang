package main

// Na prática: exercício #4
// Crie um programa que:
// atribua um valor int a uma variavel
// Demonstre este valor em decimal, binario e hexadecimal
// Desloque os bits dessa variavel 1 para esquerda e atribua a outra variavel
// Demonstre esta outra variavel em decimal, binario e hexadecimal

import (
	"fmt"
)

func main() {
	x := 45

	fmt.Printf("%v = %d, %b, %#x \n\n", x, x, x, x)

	y := x << 1

	fmt.Printf("%v = %d, %b, %#x \n\n", y, y, y, y)

}
