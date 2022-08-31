package main

import (
	"flag"
	"fmt"
)

// Flags de linha de comando são uma forma fácil de especificar opções
// para programas de linha de comando. Por exemplo, em "wc -l" o "-l" é
// uma flag de linha de comando.
// Mais em: https://en.wikipedia.org/wiki/Command-line_interface#Command-line_option

// Go fornece um pacote de suporte básico para análise de flags de linha de
// comando.
func main() {

	// A declaração básica de flags está disponível para string, inteiro, e booleano.
	// Aqui é declarado uma string flag "word" com valor padrão "foo" e uma descrição
	// curta. A função "flag.String" retorna um ponteiro string(e não um valor string).
	// Abaixo vemos como usar este ponteiro.
	wordPtr := flag.String("word", "foo", "a string")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
}
