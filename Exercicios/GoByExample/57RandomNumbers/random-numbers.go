package main

import (
	"fmt"
	"math/rand"
)

// O pacote math/rand provê geração de números pseudo-aleatórios
// Mais em: https://en.wikipedia.org/wiki/Pseudorandom_number_generator

func main() {

	// Este exemplo imprime 2 números aleatórios inteiros entre 0 e 100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// Este exemplo retorna um ponto flutuante de 0.0 à f < 1.0
	fmt.Println(rand.Float64())

	// Também é possível gerar um ponto flutuante em outra variação
	// por exemplo 5.0 <= f < 10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

}
