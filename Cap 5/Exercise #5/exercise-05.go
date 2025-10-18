package main

// Na prática: exercício #5
// Crie uma variavel de tipo string utilizando uma raw string literal
// Demonstrea

import (
	"fmt"
)

func main() {
	mensagem := ` Ola
	MUNDO`

	fmt.Printf("%v", mensagem)
}
