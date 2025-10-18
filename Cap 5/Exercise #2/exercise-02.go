package main

// Na prática: exercício #2
// Escreva expressoes utilizando os seguintes operadores, e atribua seus
// valores a variaveis
// == <= >= < >

import (
	"fmt"
)

var equal bool
var equal_and_minor bool
var equal_and_higher bool
var equal_minor bool
var equal_higher bool

func main() {
	equal = 5 == 5
	equal_and_minor = 5 <= 6
	equal_and_higher = 6 >= 5
	equal_minor = 5 < 10
	equal_higher = 6 > 100

	fmt.Printf("5 == 5 = %v \n", equal)
	fmt.Printf("5 <= 6 = %v \n", equal_and_minor)
	fmt.Printf("6 >= 5 = %v \n", equal_and_higher)
	fmt.Printf("5 < 10 = %v \n", equal_minor)
	fmt.Printf("6 > 100 = %v \n", equal_higher)
}
