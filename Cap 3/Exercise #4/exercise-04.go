package main

// Na prática: exercício #4
// Crie um tipo. O tipo subjacente deve ser int.
// Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
// Na função main:
// - Demonstre o valor da variável "x"
// - Demonstre o tipo da variável "x"
// - Atribua 42 à variável "x" utilizando o operador "="
// - Demonstre o valor da variável "x"
// Para os aventureiros: https://golang.org/ref/spec#Types
// Agora já somos quase ninjas nível 1!

import (
	"fmt"
)

type galaxya int

var x galaxya

func main() {
	fmt.Printf("Value of x: %v\n", x)
	fmt.Printf("Type of x: %T\n", x)
	x = 42
	fmt.Printf("Value of x: %v\n", x)
}
