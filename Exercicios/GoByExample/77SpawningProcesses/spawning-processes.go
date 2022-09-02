package main

import (
	"fmt"
	"io"
	"os/exec"
)

// As vezes é necessário lançar outros processos de dentro da aplicação Go.
func main() {

	// Exemplo simples de um comando que não leva argumentos e que se captura a
	// saída para imprimir na tela.
	dateCmd := exec.Command("date")

	// O método "Output" roda o comando, aguarda o término e coleta a saída
	// padrão. Se nenhum erro ocorrer "dateOut" vai conter bytes com a
	// informação da data.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// ------------

	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// ------------

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	grepCmd.Start()

	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()

	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// ------------

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
