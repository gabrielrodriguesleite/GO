package main

import (
	"fmt"
	"os"
)

// Argumentos de linha de comando são meios comuns de parametrizar a
// execução de programas. Por exemplo "go run hello.go" usa "run" e
// "hello.go" como argumentos para o programa "go".
// Mais em: https://en.wikipedia.org/wiki/Command-line_interface#Arguments

func main() {

	// os.Args fornece acesso aos argumentos crus da linha de comando.
	// O primeiro item do slice é o nome do caminho do programa, já
	// os.Args[1:] contém os argumentos do programa.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// É possível acessar os argumentos individualmente com indexação normal.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

// IMPORTANTE - para esse exemplo funcionar é necessário compilar antes:
// go build command-line-arguments.go

// Então executar incluindo argumentos:
// ./command-line-arguments a b c d

/* SAÍDA ESPERADA DESTE CÓDIGO COMPILADO E EXECUTADO COM 4 ARGUMENTOS
[./command-line-arguments a b c d]
[a b c d]
c
*/
