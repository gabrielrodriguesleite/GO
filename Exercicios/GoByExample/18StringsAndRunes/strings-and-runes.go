package main

import (
	"fmt"
	"unicode/utf8"
)

// string em go é um slice de bytes somente leitura
// go trata strings como containers de texto codificados em UTF-8
// em go strings não são feitas de "caracteres"
// o conceito de caractere é chamado de rune, um inteiro representando um ponteiro para Unicode.
// mais em: https://go.dev/blog/strings
func main() {

	// olá em thai. literals são codificados em texto UTF-8
	const s = "สวัสดี"

	// strings são equivaletes a []byte, por isso o comprimento será equivalente
	// a quantidade de bytes armazenados:
	fmt.Println("Len:", len(s)) // Len: 18

	// valores crus dos bytes para cada índice da string
	// esse for mostra os valores em hexadecimal:
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	// e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5

	// Para contar quantos runes existem em uma string podemos usar o pacote utf8
	// alguns caracteres thai são representados por multiplos ponteiros UTF-8 então
	// o resultado pode parecer estranho
	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // 6

	// um loop range manipula strings de forma especial e decodifica cada rune junto com seu offset
	// na string
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx) // U+0E2A 'ส' starts at 0 ...continua
	}

	// podemos conseguir o mesmo tipo de iteração usando a função utf8.DecodeRuneInString
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeVa, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U start at %d\n", runeVa, i)
		w = width

		// desmonstra como passar um valor rune para como argumento
		examineRune(runeVa)
	}
}

// valores entre ' aspas simples são rune literais
// podemos comparar um valor rune com um rune literal diretamente
func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}

}
