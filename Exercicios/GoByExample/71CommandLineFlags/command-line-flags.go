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

	// Aqui é declaradas as flags "numb" e "fork" usando uma abordagem similar a
	// anterior.
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// Também é possível definir uma variável e passar ela como referência para
	// que receba o valor de flag.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Depois de declarar as variáveis, uma chamada para "flag.Parse()" executa a
	// análise da linha de comando.
	flag.Parse()

	// Neste exemplo será apenas escrito na saída os valores dos argumentos.
	// Note que é necessário dereferenciar os ponteiros para conseguir acessar
	// os valores propriamente ditos.
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

// SAÍDA ESPERADA PARA ESTE CÓDIGO CONFORME A FLAG USADA:

/*
# ----------------------------------------------------------
$ go run 71CommandLineFlags/command-line-flags.go -word=opt
-numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []
# ----------------------------------------------------------
$ go run 71CommandLineFlags/command-line-flags.go -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []
# ----------------------------------------------------------
$ go run 71CommandLineFlags/command-line-flags.go -word=opt a1 a2 a3
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3]
# ----------------------------------------------------------
$ go run 71CommandLineFlags/command-line-flags.go -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]
# ----------------------------------------------------------
$ go run 71CommandLineFlags/command-line-flags.go -h
Usage of /tmp/go-build3344106437/b001/exe/command-line-flags:
  -fork
        a bool
  -numb int
        an int (default 42)
  -svar string
        a string var (default "bar")
  -word string
        a string (default "foo")
# ----------------------------------------------------------
$ go run 71CommandLineFlags/command-line-flags.go -wat
flag provided but not defined: -wat
Usage of /tmp/go-build676409743/b001/exe/command-line-flags:
  -fork
        a bool
  -numb int
        an int (default 42)
  -svar string
        a string var (default "bar")
  -word string
        a string (default "foo")
exit status 2
*/
