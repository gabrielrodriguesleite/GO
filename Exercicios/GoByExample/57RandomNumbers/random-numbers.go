package main

import (
	"fmt"
	"math/rand"
)

// O pacote math/rand provê geração de números pseudo-aleatórios
// Mais em: https://en.wikipedia.org/wiki/Pseudorandom_number_generator

func main() {

	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
}
