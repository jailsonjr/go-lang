package main

// EXERCÍCIO 01
// Utilizando o operador curto de declaração, atribua estes valores ás variáveis
// com os identificadores "x","y" e "z".
// 1. 42
// 2. "James Bond"
// 3. true
// Agora desmonstre nestas variáveis, com:
// - Uma única declaração print
// - Múltiplas declarações print

import (
	"fmt"
)

func main() {
	x, y, z := 42, "James Bond", true

	fmt.Printf("x: %v \ny: %v \nz: %v\n", x, y, z)

	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)
}
