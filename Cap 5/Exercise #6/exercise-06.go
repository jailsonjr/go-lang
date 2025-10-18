package main

// Na prática: exercício #6
// Utilizando iota, crie 4 constantes cujos valores
// sejam os proximos 4 anos

import (
	"fmt"
)

const (
	year_2026 = 2026 + iota
	year_2027
	year_2028
	year_2029
)

func main() {
	fmt.Printf("next Years: %v , %v, %v, %v", year_2026, year_2027, year_2028, year_2029)
}
