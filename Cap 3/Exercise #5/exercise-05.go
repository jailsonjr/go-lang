package main

// Na prática: exercício #5
// Utilizando a solução do exercício anterior:
// Em package-level scope, utilizando a palavra-chave var, crie uma variável com o
// identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo
// que você criou no exercício anterior.
// Na função main:
// a. Isto já deve estar feito:
// - Demonstre o valor da variável "x"
// - Demonstre o tipo da variável "x"
// - Atribua 42 à variável "x" utilizando o operador "="
// - Demonstre o valor da variável "x"
// b. Agora faça tambem:
// - Utilize conversão para transformar o tipo do valor da variável "x" em seu
// tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
// - Demonstre o valor de "y"
// - Demonstre o tipo de "y"

import (
	"fmt"
)

type galaxya int

var x galaxya
var y int

func main() {
	fmt.Printf("Value of x: %v\n", x)
	fmt.Printf("Type of x: %T\n", x)
	x = 42
	fmt.Printf("Value of x: %v\n", x)
	y = int(x)
	fmt.Printf("Value of y: %v\n", y)
	fmt.Printf("Type of y: %T\n", y)
}
