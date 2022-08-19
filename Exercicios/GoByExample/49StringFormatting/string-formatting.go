package main

import "fmt"

// go oferece um excelente suport para formatação de strings seguindo
// o modelo pringf. Neste código temos alguns exemplos.

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	// go oferece varios "verbos" de formatação desenhados para formatar
	// valorer comuns. Neste exemplo, imprimimos uma instancia da nossa
	// estrutura point:
	fmt.Printf("struct1: %v\n", p)

	fmt.Printf("struct2: %+v\n", p)

	fmt.Printf("struct3: %#v\n", p)

}
