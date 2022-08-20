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

	// -----------

	// %b para representar um binário
	fmt.Printf("bin: %b\n", 8)

	// %c para rep resentar um caractere correspondente ao inteiro
	fmt.Printf("char: %c\n", 65)

	// %x apresenta codificação hexadecimal
	fmt.Printf("hex: %x\n", 10)

	// %f imprime uma das diversas representações para ponto flutuante
	fmt.Printf("float1: %f\n", 65.32)

	// %e e %E apresentam uma formatação de notação científica
	fmt.Printf("float2: %e\n", 65.32)

	// um com expoênte em maiúsculo e outra em minúsculo
	fmt.Printf("float3: %E\n", 65.32)

	// %s para formatação mais simples para string
	fmt.Printf("str1: %s\n", "\"string\"")

	// %q é para imprimir a barra com aspas duplas como aqui no código fonte
	fmt.Printf("str2: %q\n", "\"string\"")

	// %x também apresenta um caractere como um byte hexadecimal
	fmt.Printf("str3: %x\n", "hex this")

	// %p representa um ponteiro (valor na memória em hexa)
	fmt.Printf("pointer: %p\n", &p)

	// %#verbo o # (inteiro) faz o controle de largura e a precisão de caracteres ao
	// representar um número. Por padrão o resultado vai ser justificado a direita
	// e preenchido com espaços.
	fmt.Printf("width1: |%6d|%6d|\n", 648, 45)

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 64.8, 4.053)

	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 68.0, 0.457)

	fmt.Printf("width4: |%6s|%6s|\n", "xoo", "k")

	fmt.Printf("width5: |%-6s|%-6s|\n", "ate", "g")

}
