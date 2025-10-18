package main

// Na prática: exercício #3
// Crie constantes tipadas e nao tipadas
// Demonstre seus valores

import (
	"fmt"
)

const (
	x int  = 42
	y      = "James Bond"
	z bool = true
)

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
