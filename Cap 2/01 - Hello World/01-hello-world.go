package main

import (
	"fmt"
)

func main() {
	_, erros := fmt.Println("Hello World", "Oi galera!", 100)
	fmt.Println(erros)

	x := 16
	w := 16.6
	y := "strings"
	z := true

	x = 10

	fmt.Println(x, y, z, w)
}
