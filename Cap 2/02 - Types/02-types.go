package main

import (
	"fmt"
)

type hotdog int

var z hotdog = 10

func main() {
	x := 16
	y := "texto"

	fmt.Printf("X: %v, %T \n", x, x)
	fmt.Printf("Y: %v, %T \n", y, y)
	fmt.Printf("Z: %v, %T \n", z, z)
}
