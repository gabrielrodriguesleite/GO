package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Line filter é um tipo de programa comum que lê a entrada do stdin,
// processa a entrada e então imprime o resultado para o stdout.
// Exemplos desses tipos são grep e sed

// Neste exemplo implementa-se um line filter em go que escreve a
// versão em letras maiúsculas de todo texto inserido.
// É possível implementar todo tipo de padrão de formatação.
func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
