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

	// Envolvendo os.Stdin que é unbuffered com um scanner buferizado
	// nos dá a conveniência do método Scan que avança o scanner para
	// o próvimo token; o qual é a próxima linha no scanner padrão.
	scanner := bufio.NewScanner(os.Stdin)

	// O método Text retorna o token atual, neste caso a próxima linha
	// da entrada.
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		// Escrevendo a versão capitalizada.
		fmt.Println(ucl)
	}

	// Verifica por erros durante o Scan. EOF (fim de arquivo) é esperado
	// então não é reportado como erro por Scan.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

/* SAÍDA ESPERADA
# echo 'hello' > /tmp/lines
# echo 'filter' >> /tmp/lines
# cat /tmp/lines | go run line-filters.go
HELLO
FILTER
# também funciona com a entrada do teclado se executado com
go run line-filters.go
*/
