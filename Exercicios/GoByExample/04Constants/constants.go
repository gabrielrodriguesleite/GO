package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 500000000

	// expressões performam aritmetica com precisão arbitraria
	const d = 3e20 / n
	fmt.Println(d)

	// numeros constantes não possuem tipo até que um tipo seja dado como por exemplo com uma conversão explicita
	fmt.Println(int64(d))

	// neste caso Sin expera um float64 então a constante receberá o tipo necessário para o contexto em questão
	fmt.Println(math.Sin(n))
}
