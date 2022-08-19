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

	// Para valores que são estruturas o %+v vai incluir os nomes dos campos
	fmt.Printf("struct2: %+v\n", p)

	// %#v imprime a sintaxe Go que representa o valor.
	// ex: o código fonte que vai produzir o valor
	fmt.Printf("struct3: %#v\n", p)

	// -----------

	// Para imprimir o tipo de um valor podemos usar %T
	fmt.Printf("type: %T\n", p)

	// Para imprimir um booleano %t
	fmt.Printf("bool: %t\n", true)

	// Existem diversas formas de imprimir inteiros
	// %d por padrão, formatação com base 10
	fmt.Printf("int: %d\n", 123)
}
