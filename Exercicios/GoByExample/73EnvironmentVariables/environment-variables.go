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

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

}
