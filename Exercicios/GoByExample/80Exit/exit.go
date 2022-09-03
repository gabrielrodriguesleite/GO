package main

import (
	"fmt"
	"os"
)

// Use "os.Exit" para sair imediatamente com um dado estado.

func main() {

	// "defer"s não vão ser executados quando sair por "os.Exit", então
	// este "fmt.Println" nunca será chamado.
	defer fmt.Println("!")

	// Sair com estado 3
	os.Exit(3)

}
