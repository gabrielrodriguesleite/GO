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

// Note que diferente de "C" por exemplo, Go não usa um inteiro como valor
// de retorno de "main" para indicar o estado de saída.
// Se é necessario sair com um estado diferente de zero então "os.Exit"
// é a solução.

// SAÍDA ESPERADA PARA ESTE CÓDIGO:

// Executando com "go run exit.go" a saída será capturada por "go" e impressa

/* go run 80Exit/exit.go
exit status 3
*/

// "echo $?" pode ser usado para mostrar o esdado do processo anterior

/*go build exit.go
./exit
echo $?
3
*/
