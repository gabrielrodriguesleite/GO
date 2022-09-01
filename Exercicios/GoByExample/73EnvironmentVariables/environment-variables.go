package main

import (
	"fmt"
	"os"
)

// Variáveis de ambiente são mecanismos universais para transmitir
// informações de configuração para programas Unix.
// Mais em: https://en.wikipedia.org/wiki/Environment_variable
// Também: https://www.12factor.net/config
// Neste exemplo as variáveis de ambiente serão listadas, escritas e lidas.

func main() {

	// Para definir um par de chave/valor se usa "os.Setenv".
	// Para ler um valor de uma chave se usa "os.Getenv".
	// O valor retornado será uma string vazia caso a chave não esteja
	// presente no ambiente.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

}
